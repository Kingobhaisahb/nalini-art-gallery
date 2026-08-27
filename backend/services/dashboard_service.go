package services

import (
	"github.com/Kingobhaisahb/nalini-art-gallery/dto"
	"github.com/Kingobhaisahb/nalini-art-gallery/repositories"
)

type DashboardService struct {
	DashboardRepo repositories.DashboardRepository
}

func (s *DashboardService) GetAdminDashboard() (*dto.AdminDashboardResponse, error) {

	totalPaintings, err := s.DashboardRepo.GetTotalPaintings()
	if err != nil {
		return nil, err
	}

	availablePaintings, err := s.DashboardRepo.GetAvailablePaintings()
	if err != nil {
		return nil, err
	}

	soldPaintings, err := s.DashboardRepo.GetSoldPaintings()
	if err != nil {
		return nil, err
	}

	featuredPaintings, err := s.DashboardRepo.GetFeaturedPaintings()
	if err != nil {
		return nil, err
	}

	totalOrders, err := s.DashboardRepo.GetTotalOrders()
	if err != nil {
		return nil, err
	}

	pendingOrders, err := s.DashboardRepo.GetOrdersByStatus("PENDING")
	if err != nil {
		return nil, err
	}

	confirmedOrders, err := s.DashboardRepo.GetOrdersByStatus("CONFIRMED")
	if err != nil {
		return nil, err
	}

	shippedOrders, err := s.DashboardRepo.GetOrdersByStatus("SHIPPED")
	if err != nil {
		return nil, err
	}

	deliveredOrders, err := s.DashboardRepo.GetOrdersByStatus("DELIVERED")
	if err != nil {
		return nil, err
	}

	cancelledOrders, err := s.DashboardRepo.GetOrdersByStatus("CANCELLED")
	if err != nil {
		return nil, err
	}

	totalCustomers, err := s.DashboardRepo.GetTotalCustomers()
	if err != nil {
		return nil, err
	}

	totalRevenue, err := s.DashboardRepo.GetTotalRevenue()
	if err != nil {
		return nil, err
	}

	return &dto.AdminDashboardResponse{
		TotalPaintings:     totalPaintings,
		AvailablePaintings: availablePaintings,
		SoldPaintings:      soldPaintings,
		FeaturedPaintings:  featuredPaintings,

		TotalOrders:     totalOrders,
		PendingOrders:   pendingOrders,
		ConfirmedOrders: confirmedOrders,
		ShippedOrders:   shippedOrders,
		DeliveredOrders: deliveredOrders,
		CancelledOrders: cancelledOrders,

		TotalCustomers: totalCustomers,

		TotalRevenue: totalRevenue,
	}, nil
}