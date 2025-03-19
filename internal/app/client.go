package app

import (
	"fmt"
	"time"

	"github.com/andrersp/favorites/internal/domain/dto"
	"github.com/andrersp/favorites/internal/domain/entity"
	"github.com/andrersp/favorites/internal/domain/ports"
	genericresponse "github.com/andrersp/favorites/pkg/generic-response"
)

var (
	errorDuplicateEmail = genericresponse.NewErrorResponse("duplicate email", "email already exists")
)

type ClientApp interface {
	Create(request dto.ClientRequestDTO) error
	Update(clientID uint, request dto.ClientRequestDTO) error
	List() ([]dto.ClientResumeResponseDTO, error)
	Detail(clientID uint) (dto.ClientDetailResponseDTO, error)
	Delete(clientID uint) error
}

type clientApp struct {
	clientRepository  ports.ClientRepository
	productRepository ports.ProductRepository
	cacheRepository   ports.CacheRepository
}

func NewClientApp(
	clientRepository ports.ClientRepository,
	productRepository ports.ProductRepository,
	cacheRepository ports.CacheRepository,
) ClientApp {
	app := new(clientApp)
	app.clientRepository = clientRepository
	app.productRepository = productRepository
	app.cacheRepository = cacheRepository

	return app
}

// Delete implements ClientApp.
func (c *clientApp) Delete(clientID uint) error {
	if err := c.clientRepository.Delete(clientID); err != nil {
		return genericresponse.NewErrorResponse("Error on delete client", err.Error())
	}

	return nil
}

// Update implements ClientApp.
func (c *clientApp) Update(clientID uint, request dto.ClientRequestDTO) error {
	existClient, err := c.clientRepository.FindByID(clientID)
	if err != nil {
		return err
	}

	clientByEmail, err := c.clientRepository.FindByEmail(request.Email)

	if err == nil && clientByEmail.ID != existClient.ID {
		return errorDuplicateEmail
	}

	existClient.Email = request.Email
	existClient.Name = request.Name

	return c.clientRepository.Update(existClient)
}

// Detail implements ClientApp.
func (c *clientApp) Detail(clientID uint) (dto.ClientDetailResponseDTO, error) {
	client, err := c.clientRepository.FindByID(clientID)
	if err != nil {
		return dto.ClientDetailResponseDTO{}, err
	}

	clientDTO := dto.ClientDetailFromEntity(client)

	for index := range client.Favorites {
		product, err := c.getProductByID(client.Favorites[index].ProductID)
		if err != nil {
			return dto.ClientDetailResponseDTO{}, err
		}

		clientDTO.AddProduct(product)
	}

	return clientDTO, nil
}

// Create implements ClientApp.
func (c *clientApp) Create(request dto.ClientRequestDTO) error {
	clientEntity := request.ToEntity()

	if _, err := c.findDuplicateUser(clientEntity); err != nil {
		return err
	}

	if err := c.clientRepository.Save(clientEntity); err != nil {
		return err
	}

	return nil
}

// List implements ClientApp.
func (c *clientApp) List() ([]dto.ClientResumeResponseDTO, error) {
	clients, err := c.clientRepository.List()
	if err != nil {
		return nil, err
	}

	return dto.ClientListFromEntity(clients), nil
}

func (c *clientApp) findDuplicateUser(client entity.Client) (entity.Client, error) {
	client, err := c.clientRepository.FindByEmail(client.Email)
	if err == nil {
		return entity.Client{}, errorDuplicateEmail
	}

	return client, nil
}

func (f *clientApp) getProductByID(productID uint) (dto.ProductResponseDTO, error) {
	var responseDTO dto.ProductResponseDTO

	productKey := fmt.Sprintf("product-detail-%d", productID)

	if err := f.cacheRepository.Get(productKey, &responseDTO); err == nil {
		return responseDTO, nil
	}

	productEntity, err := f.productRepository.Detail(productID)
	if err != nil {
		return dto.ProductResponseDTO{}, err
	}

	responseDTO = dto.ProductEntityToProductResponseDTO(productEntity)

	_ = f.cacheRepository.Set(productKey, responseDTO, productTimeExpiration*time.Minute)

	return responseDTO, nil
}
