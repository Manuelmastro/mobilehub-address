package services

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/Manuelmastro/mobilehub-address/pkg/db"
	"github.com/Manuelmastro/mobilehub-address/pkg/models"
	"github.com/Manuelmastro/mobilehub-address/pkg/pb"

	"gorm.io/gorm"
)

type AddressServiceServer struct {
	pb.UnimplementedAddressServiceServer
	H db.Handler
}

// Add a new address
func (s *AddressServiceServer) AddAddress(ctx context.Context, req *pb.AddAddressRequest) (*pb.AddAddressResponse, error) {
	address := models.Address{
		UserID:     req.Address.UserId,
		Country:    req.Address.Country,
		State:      req.Address.State,
		District:   req.Address.District,
		StreetName: req.Address.StreetName,
		PinCode:    req.Address.PinCode,
		Phone:      req.Address.Phone,
	}

	if err := s.H.DB.Create(&address).Error; err != nil {
		return nil, errors.New("failed to add address")
	}

	return &pb.AddAddressResponse{
		Message: "Address added successfully",
		Id:      fmt.Sprint(address.ID),
	}, nil
}

// List all addresses of a user
func (s *AddressServiceServer) ListAddresses(ctx context.Context, req *pb.ListAddressesRequest) (*pb.ListAddressesResponse, error) {
	var addresses []models.Address
	if err := s.H.DB.Where("user_id = ?", req.UserId).Find(&addresses).Error; err != nil {
		return nil, errors.New("failed to fetch addresses")
	}

	var response []*pb.Address
	for _, address := range addresses {
		response = append(response, &pb.Address{
			Id:         fmt.Sprint(address.ID),
			UserId:     address.UserID,
			Country:    address.Country,
			State:      address.State,
			District:   address.District,
			StreetName: address.StreetName,
			PinCode:    address.PinCode,
			Phone:      address.Phone,
		})
	}

	return &pb.ListAddressesResponse{Addresses: response}, nil
}

// Get a specific address by ID
func (s *AddressServiceServer) GetAddress(ctx context.Context, req *pb.GetAddressRequest) (*pb.GetAddressResponse, error) {
	var address models.Address
	if err := s.H.DB.Where("id = ?", req.Id).First(&address).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("address not found")
		}
		return nil, errors.New("failed to fetch address")
	}

	return &pb.GetAddressResponse{
		Address: &pb.Address{
			Id:         fmt.Sprint(address.ID),
			UserId:     address.UserID,
			Country:    address.Country,
			State:      address.State,
			District:   address.District,
			StreetName: address.StreetName,
			PinCode:    address.PinCode,
			Phone:      address.Phone,
		},
	}, nil
}

// Edit an existing address
func (s *AddressServiceServer) EditAddress(ctx context.Context, req *pb.EditAddressRequest) (*pb.EditAddressResponse, error) {
	addressID, err := strconv.Atoi(req.Address.Id)
	if err != nil {
		return nil, errors.New("invalid address ID")
	}

	var address models.Address
	if err := s.H.DB.First(&address, addressID).Error; err != nil {
		return nil, errors.New("address not found")
	}

	address.Country = req.Address.Country
	address.State = req.Address.State
	address.District = req.Address.District
	address.StreetName = req.Address.StreetName
	address.PinCode = req.Address.PinCode
	address.Phone = req.Address.Phone

	if err := s.H.DB.Save(&address).Error; err != nil {
		return nil, errors.New("failed to update address")
	}

	return &pb.EditAddressResponse{
		Message: "Address updated successfully",
	}, nil
}

// Delete an address
func (s *AddressServiceServer) DeleteAddress(ctx context.Context, req *pb.DeleteAddressRequest) (*pb.DeleteAddressResponse, error) {
	addressID, err := strconv.Atoi(req.Id)
	if err != nil {
		return nil, errors.New("invalid address ID")
	}

	if err := s.H.DB.Delete(&models.Address{}, addressID).Error; err != nil {
		return nil, errors.New("failed to delete address")
	}

	return &pb.DeleteAddressResponse{
		Message: "Address deleted successfully",
	}, nil
}
