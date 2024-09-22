# GetOrders200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Orders** | Pointer to [**[]GetOrderSchema**](GetOrderSchema.md) |  | [optional] 

## Methods

### NewGetOrders200Response

`func NewGetOrders200Response() *GetOrders200Response`

NewGetOrders200Response instantiates a new GetOrders200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrders200ResponseWithDefaults

`func NewGetOrders200ResponseWithDefaults() *GetOrders200Response`

NewGetOrders200ResponseWithDefaults instantiates a new GetOrders200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrders

`func (o *GetOrders200Response) GetOrders() []GetOrderSchema`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *GetOrders200Response) GetOrdersOk() (*[]GetOrderSchema, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *GetOrders200Response) SetOrders(v []GetOrderSchema)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *GetOrders200Response) HasOrders() bool`

HasOrders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


