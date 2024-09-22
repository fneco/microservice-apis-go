# OrderItemSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | **string** |  | 
**Size** | **string** |  | 
**Quantity** | Pointer to **int64** |  | [optional] [default to 1]

## Methods

### NewOrderItemSchema

`func NewOrderItemSchema(product string, size string, ) *OrderItemSchema`

NewOrderItemSchema instantiates a new OrderItemSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderItemSchemaWithDefaults

`func NewOrderItemSchemaWithDefaults() *OrderItemSchema`

NewOrderItemSchemaWithDefaults instantiates a new OrderItemSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *OrderItemSchema) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *OrderItemSchema) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *OrderItemSchema) SetProduct(v string)`

SetProduct sets Product field to given value.


### GetSize

`func (o *OrderItemSchema) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *OrderItemSchema) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *OrderItemSchema) SetSize(v string)`

SetSize sets Size field to given value.


### GetQuantity

`func (o *OrderItemSchema) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *OrderItemSchema) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *OrderItemSchema) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *OrderItemSchema) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


