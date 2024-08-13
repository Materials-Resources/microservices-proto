// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: proto/catalog/v1/catalog.proto

package catalog

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	CatalogService_ListProducts_FullMethodName         = "/catalog.v1.CatalogService/ListProducts"
	CatalogService_CreateProduct_FullMethodName        = "/catalog.v1.CatalogService/CreateProduct"
	CatalogService_GetProduct_FullMethodName           = "/catalog.v1.CatalogService/GetProduct"
	CatalogService_UpdateProduct_FullMethodName        = "/catalog.v1.CatalogService/UpdateProduct"
	CatalogService_DeleteProduct_FullMethodName        = "/catalog.v1.CatalogService/DeleteProduct"
	CatalogService_ListGroups_FullMethodName           = "/catalog.v1.CatalogService/ListGroups"
	CatalogService_CreateGroup_FullMethodName          = "/catalog.v1.CatalogService/CreateGroup"
	CatalogService_GetGroup_FullMethodName             = "/catalog.v1.CatalogService/GetGroup"
	CatalogService_UpdateGroup_FullMethodName          = "/catalog.v1.CatalogService/UpdateGroup"
	CatalogService_DeleteGroup_FullMethodName          = "/catalog.v1.CatalogService/DeleteGroup"
	CatalogService_ListSuppliers_FullMethodName        = "/catalog.v1.CatalogService/ListSuppliers"
	CatalogService_GetSupplier_FullMethodName          = "/catalog.v1.CatalogService/GetSupplier"
	CatalogService_CreateSupplier_FullMethodName       = "/catalog.v1.CatalogService/CreateSupplier"
	CatalogService_UpdateSupplier_FullMethodName       = "/catalog.v1.CatalogService/UpdateSupplier"
	CatalogService_SetPrimarySupplier_FullMethodName   = "/catalog.v1.CatalogService/SetPrimarySupplier"
	CatalogService_GetProductBySupplier_FullMethodName = "/catalog.v1.CatalogService/GetProductBySupplier"
)

// CatalogServiceClient is the client API for CatalogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CatalogServiceClient interface {
	// ListProducts returns a List of products
	ListProducts(ctx context.Context, in *ListProductsRequest, opts ...grpc.CallOption) (*ListProductsResponse, error)
	// CreateProduct creates a product
	CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*CreateProductResponse, error)
	// GetProduct returns a product
	GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*GetProductResponse, error)
	// UpdateProduct updates a product
	UpdateProduct(ctx context.Context, in *UpdateProductRequest, opts ...grpc.CallOption) (*UpdateProductResponse, error)
	// DeleteProduct deletes a product
	DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*DeleteProductResponse, error)
	// ListGroups returns a List of product groups
	ListGroups(ctx context.Context, in *ListGroupRequest, opts ...grpc.CallOption) (*ListGroupResponse, error)
	// CreateGroup creates a new product group
	CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*CreateGroupResponse, error)
	// ReadGroup returns a product group with products
	GetGroup(ctx context.Context, in *GetGroupRequest, opts ...grpc.CallOption) (*GetGroupResponse, error)
	// UpdateGroup updates a product group
	UpdateGroup(ctx context.Context, in *UpdateGroupRequest, opts ...grpc.CallOption) (*UpdateGroupResponse, error)
	// DeleteGroup deletes a product group
	DeleteGroup(ctx context.Context, in *DeleteGroupRequest, opts ...grpc.CallOption) (*DeleteGroupResponse, error)
	// ListSuppliers returns a List of suppliers for a product
	ListSuppliers(ctx context.Context, in *ListSuppliersRequest, opts ...grpc.CallOption) (*ListSuppliersResponse, error)
	// GetSupplier returns the supplier of a product
	GetSupplier(ctx context.Context, in *GetSupplierRequest, opts ...grpc.CallOption) (*GetSupplierResponse, error)
	// CreateSupplier creates a supplier for a product
	CreateSupplier(ctx context.Context, in *CreateSupplierRequest, opts ...grpc.CallOption) (*CreateSupplierResponse, error)
	// UpdateSupplier updates a supplier for a product
	UpdateSupplier(ctx context.Context, in *UpdateSupplierRequest, opts ...grpc.CallOption) (*UpdateSupplierResponse, error)
	// SetPrimarySupplier sets the primary supplier for a product
	SetPrimarySupplier(ctx context.Context, in *SetPrimarySupplierRequest, opts ...grpc.CallOption) (*SetPrimarySupplierResponse, error)
	// GetProductBySupplier returns a product by supplier identifiers
	GetProductBySupplier(ctx context.Context, in *GetProductBySupplierRequest, opts ...grpc.CallOption) (*GetProductBySupplierResponse, error)
}

type catalogServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCatalogServiceClient(cc grpc.ClientConnInterface) CatalogServiceClient {
	return &catalogServiceClient{cc}
}

func (c *catalogServiceClient) ListProducts(ctx context.Context, in *ListProductsRequest, opts ...grpc.CallOption) (*ListProductsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListProductsResponse)
	err := c.cc.Invoke(ctx, CatalogService_ListProducts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*CreateProductResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateProductResponse)
	err := c.cc.Invoke(ctx, CatalogService_CreateProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*GetProductResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetProductResponse)
	err := c.cc.Invoke(ctx, CatalogService_GetProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) UpdateProduct(ctx context.Context, in *UpdateProductRequest, opts ...grpc.CallOption) (*UpdateProductResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateProductResponse)
	err := c.cc.Invoke(ctx, CatalogService_UpdateProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*DeleteProductResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteProductResponse)
	err := c.cc.Invoke(ctx, CatalogService_DeleteProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) ListGroups(ctx context.Context, in *ListGroupRequest, opts ...grpc.CallOption) (*ListGroupResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListGroupResponse)
	err := c.cc.Invoke(ctx, CatalogService_ListGroups_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*CreateGroupResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateGroupResponse)
	err := c.cc.Invoke(ctx, CatalogService_CreateGroup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) GetGroup(ctx context.Context, in *GetGroupRequest, opts ...grpc.CallOption) (*GetGroupResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetGroupResponse)
	err := c.cc.Invoke(ctx, CatalogService_GetGroup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) UpdateGroup(ctx context.Context, in *UpdateGroupRequest, opts ...grpc.CallOption) (*UpdateGroupResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateGroupResponse)
	err := c.cc.Invoke(ctx, CatalogService_UpdateGroup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) DeleteGroup(ctx context.Context, in *DeleteGroupRequest, opts ...grpc.CallOption) (*DeleteGroupResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteGroupResponse)
	err := c.cc.Invoke(ctx, CatalogService_DeleteGroup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) ListSuppliers(ctx context.Context, in *ListSuppliersRequest, opts ...grpc.CallOption) (*ListSuppliersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListSuppliersResponse)
	err := c.cc.Invoke(ctx, CatalogService_ListSuppliers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) GetSupplier(ctx context.Context, in *GetSupplierRequest, opts ...grpc.CallOption) (*GetSupplierResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSupplierResponse)
	err := c.cc.Invoke(ctx, CatalogService_GetSupplier_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) CreateSupplier(ctx context.Context, in *CreateSupplierRequest, opts ...grpc.CallOption) (*CreateSupplierResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateSupplierResponse)
	err := c.cc.Invoke(ctx, CatalogService_CreateSupplier_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) UpdateSupplier(ctx context.Context, in *UpdateSupplierRequest, opts ...grpc.CallOption) (*UpdateSupplierResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateSupplierResponse)
	err := c.cc.Invoke(ctx, CatalogService_UpdateSupplier_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) SetPrimarySupplier(ctx context.Context, in *SetPrimarySupplierRequest, opts ...grpc.CallOption) (*SetPrimarySupplierResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetPrimarySupplierResponse)
	err := c.cc.Invoke(ctx, CatalogService_SetPrimarySupplier_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) GetProductBySupplier(ctx context.Context, in *GetProductBySupplierRequest, opts ...grpc.CallOption) (*GetProductBySupplierResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetProductBySupplierResponse)
	err := c.cc.Invoke(ctx, CatalogService_GetProductBySupplier_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CatalogServiceServer is the server API for CatalogService service.
// All implementations should embed UnimplementedCatalogServiceServer
// for forward compatibility.
type CatalogServiceServer interface {
	// ListProducts returns a List of products
	ListProducts(context.Context, *ListProductsRequest) (*ListProductsResponse, error)
	// CreateProduct creates a product
	CreateProduct(context.Context, *CreateProductRequest) (*CreateProductResponse, error)
	// GetProduct returns a product
	GetProduct(context.Context, *GetProductRequest) (*GetProductResponse, error)
	// UpdateProduct updates a product
	UpdateProduct(context.Context, *UpdateProductRequest) (*UpdateProductResponse, error)
	// DeleteProduct deletes a product
	DeleteProduct(context.Context, *DeleteProductRequest) (*DeleteProductResponse, error)
	// ListGroups returns a List of product groups
	ListGroups(context.Context, *ListGroupRequest) (*ListGroupResponse, error)
	// CreateGroup creates a new product group
	CreateGroup(context.Context, *CreateGroupRequest) (*CreateGroupResponse, error)
	// ReadGroup returns a product group with products
	GetGroup(context.Context, *GetGroupRequest) (*GetGroupResponse, error)
	// UpdateGroup updates a product group
	UpdateGroup(context.Context, *UpdateGroupRequest) (*UpdateGroupResponse, error)
	// DeleteGroup deletes a product group
	DeleteGroup(context.Context, *DeleteGroupRequest) (*DeleteGroupResponse, error)
	// ListSuppliers returns a List of suppliers for a product
	ListSuppliers(context.Context, *ListSuppliersRequest) (*ListSuppliersResponse, error)
	// GetSupplier returns the supplier of a product
	GetSupplier(context.Context, *GetSupplierRequest) (*GetSupplierResponse, error)
	// CreateSupplier creates a supplier for a product
	CreateSupplier(context.Context, *CreateSupplierRequest) (*CreateSupplierResponse, error)
	// UpdateSupplier updates a supplier for a product
	UpdateSupplier(context.Context, *UpdateSupplierRequest) (*UpdateSupplierResponse, error)
	// SetPrimarySupplier sets the primary supplier for a product
	SetPrimarySupplier(context.Context, *SetPrimarySupplierRequest) (*SetPrimarySupplierResponse, error)
	// GetProductBySupplier returns a product by supplier identifiers
	GetProductBySupplier(context.Context, *GetProductBySupplierRequest) (*GetProductBySupplierResponse, error)
}

// UnimplementedCatalogServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCatalogServiceServer struct{}

func (UnimplementedCatalogServiceServer) ListProducts(context.Context, *ListProductsRequest) (*ListProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProducts not implemented")
}
func (UnimplementedCatalogServiceServer) CreateProduct(context.Context, *CreateProductRequest) (*CreateProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProduct not implemented")
}
func (UnimplementedCatalogServiceServer) GetProduct(context.Context, *GetProductRequest) (*GetProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (UnimplementedCatalogServiceServer) UpdateProduct(context.Context, *UpdateProductRequest) (*UpdateProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProduct not implemented")
}
func (UnimplementedCatalogServiceServer) DeleteProduct(context.Context, *DeleteProductRequest) (*DeleteProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProduct not implemented")
}
func (UnimplementedCatalogServiceServer) ListGroups(context.Context, *ListGroupRequest) (*ListGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGroups not implemented")
}
func (UnimplementedCatalogServiceServer) CreateGroup(context.Context, *CreateGroupRequest) (*CreateGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroup not implemented")
}
func (UnimplementedCatalogServiceServer) GetGroup(context.Context, *GetGroupRequest) (*GetGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroup not implemented")
}
func (UnimplementedCatalogServiceServer) UpdateGroup(context.Context, *UpdateGroupRequest) (*UpdateGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGroup not implemented")
}
func (UnimplementedCatalogServiceServer) DeleteGroup(context.Context, *DeleteGroupRequest) (*DeleteGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGroup not implemented")
}
func (UnimplementedCatalogServiceServer) ListSuppliers(context.Context, *ListSuppliersRequest) (*ListSuppliersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSuppliers not implemented")
}
func (UnimplementedCatalogServiceServer) GetSupplier(context.Context, *GetSupplierRequest) (*GetSupplierResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSupplier not implemented")
}
func (UnimplementedCatalogServiceServer) CreateSupplier(context.Context, *CreateSupplierRequest) (*CreateSupplierResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSupplier not implemented")
}
func (UnimplementedCatalogServiceServer) UpdateSupplier(context.Context, *UpdateSupplierRequest) (*UpdateSupplierResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSupplier not implemented")
}
func (UnimplementedCatalogServiceServer) SetPrimarySupplier(context.Context, *SetPrimarySupplierRequest) (*SetPrimarySupplierResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPrimarySupplier not implemented")
}
func (UnimplementedCatalogServiceServer) GetProductBySupplier(context.Context, *GetProductBySupplierRequest) (*GetProductBySupplierResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductBySupplier not implemented")
}
func (UnimplementedCatalogServiceServer) testEmbeddedByValue() {}

// UnsafeCatalogServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CatalogServiceServer will
// result in compilation errors.
type UnsafeCatalogServiceServer interface {
	mustEmbedUnimplementedCatalogServiceServer()
}

func RegisterCatalogServiceServer(s grpc.ServiceRegistrar, srv CatalogServiceServer) {
	// If the following call pancis, it indicates UnimplementedCatalogServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CatalogService_ServiceDesc, srv)
}

func _CatalogService_ListProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).ListProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatalogService_ListProducts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).ListProducts(ctx, req.(*ListProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_CreateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).CreateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatalogService_CreateProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).CreateProduct(ctx, req.(*CreateProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatalogService_GetProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).GetProduct(ctx, req.(*GetProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_UpdateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).UpdateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatalogService_UpdateProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).UpdateProduct(ctx, req.(*UpdateProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_DeleteProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).DeleteProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatalogService_DeleteProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).DeleteProduct(ctx, req.(*DeleteProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_ListGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).ListGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatalogService_ListGroups_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).ListGroups(ctx, req.(*ListGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_CreateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).CreateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatalogService_CreateGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).CreateGroup(ctx, req.(*CreateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_GetGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).GetGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatalogService_GetGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).GetGroup(ctx, req.(*GetGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_UpdateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).UpdateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatalogService_UpdateGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).UpdateGroup(ctx, req.(*UpdateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_DeleteGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).DeleteGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatalogService_DeleteGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).DeleteGroup(ctx, req.(*DeleteGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_ListSuppliers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSuppliersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).ListSuppliers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatalogService_ListSuppliers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).ListSuppliers(ctx, req.(*ListSuppliersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_GetSupplier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSupplierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).GetSupplier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatalogService_GetSupplier_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).GetSupplier(ctx, req.(*GetSupplierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_CreateSupplier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSupplierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).CreateSupplier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatalogService_CreateSupplier_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).CreateSupplier(ctx, req.(*CreateSupplierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_UpdateSupplier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSupplierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).UpdateSupplier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatalogService_UpdateSupplier_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).UpdateSupplier(ctx, req.(*UpdateSupplierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_SetPrimarySupplier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetPrimarySupplierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).SetPrimarySupplier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatalogService_SetPrimarySupplier_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).SetPrimarySupplier(ctx, req.(*SetPrimarySupplierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_GetProductBySupplier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductBySupplierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).GetProductBySupplier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatalogService_GetProductBySupplier_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).GetProductBySupplier(ctx, req.(*GetProductBySupplierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CatalogService_ServiceDesc is the grpc.ServiceDesc for CatalogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CatalogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "catalog.v1.CatalogService",
	HandlerType: (*CatalogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListProducts",
			Handler:    _CatalogService_ListProducts_Handler,
		},
		{
			MethodName: "CreateProduct",
			Handler:    _CatalogService_CreateProduct_Handler,
		},
		{
			MethodName: "GetProduct",
			Handler:    _CatalogService_GetProduct_Handler,
		},
		{
			MethodName: "UpdateProduct",
			Handler:    _CatalogService_UpdateProduct_Handler,
		},
		{
			MethodName: "DeleteProduct",
			Handler:    _CatalogService_DeleteProduct_Handler,
		},
		{
			MethodName: "ListGroups",
			Handler:    _CatalogService_ListGroups_Handler,
		},
		{
			MethodName: "CreateGroup",
			Handler:    _CatalogService_CreateGroup_Handler,
		},
		{
			MethodName: "GetGroup",
			Handler:    _CatalogService_GetGroup_Handler,
		},
		{
			MethodName: "UpdateGroup",
			Handler:    _CatalogService_UpdateGroup_Handler,
		},
		{
			MethodName: "DeleteGroup",
			Handler:    _CatalogService_DeleteGroup_Handler,
		},
		{
			MethodName: "ListSuppliers",
			Handler:    _CatalogService_ListSuppliers_Handler,
		},
		{
			MethodName: "GetSupplier",
			Handler:    _CatalogService_GetSupplier_Handler,
		},
		{
			MethodName: "CreateSupplier",
			Handler:    _CatalogService_CreateSupplier_Handler,
		},
		{
			MethodName: "UpdateSupplier",
			Handler:    _CatalogService_UpdateSupplier_Handler,
		},
		{
			MethodName: "SetPrimarySupplier",
			Handler:    _CatalogService_SetPrimarySupplier_Handler,
		},
		{
			MethodName: "GetProductBySupplier",
			Handler:    _CatalogService_GetProductBySupplier_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/catalog/v1/catalog.proto",
}