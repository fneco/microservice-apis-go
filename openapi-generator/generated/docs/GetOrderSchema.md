# GetOrderSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Created** | **time.Time** |  | 
**Status** | **string** |  | 
**Order** | [**[]OrderItemSchema**](OrderItemSchema.md) |  | 

## Methods

### NewGetOrderSchema

`func NewGetOrderSchema(id string, created time.Time, status string, order []OrderItemSchema, ) *GetOrderSchema`

NewGetOrderSchema instantiates a new GetOrderSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrderSchemaWithDefaults

`func NewGetOrderSchemaWithDefaults() *GetOrderSchema`

NewGetOrderSchemaWithDefaults instantiates a new GetOrderSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetOrderSchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetOrderSchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetOrderSchema) SetId(v string)`

SetId sets Id field to given value.


### GetCreated

`func (o *GetOrderSchema) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GetOrderSchema) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GetOrderSchema) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetStatus

`func (o *GetOrderSchema) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetOrderSchema) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetOrderSchema) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetOrder

`func (o *GetOrderSchema) GetOrder() []OrderItemSchema`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *GetOrderSchema) GetOrderOk() (*[]OrderItemSchema, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *GetOrderSchema) SetOrder(v []OrderItemSchema)`

SetOrder sets Order field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


