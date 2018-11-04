GeoOrderTest .
==============
Used for customer clients, then can place order between different places, take order, and list orders. Use Google Maps API to get the distance for the order.

The markdown API doc does not has responses which is a bug. The http server API doc is better.

**Version:** 0.1

### /orders
---
##### ***GET***
**Summary:** List orders.

**Parameters**

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| page | path | Page index from 1 | Yes | string |
| limit | path | Order numbers per page, between 1 to 50 | Yes | string |

**Responses**

| Code | Description |
| ---- | ----------- |
| 200 |  |
| default |  |

##### ***POST***
**Summary:** Create an order with an origin and a destination.

**Description:** Request:
default: placeRequest

**Parameters**

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| PlaceRequest | body |  | Yes | [placeRequest](#placerequest) |

**Responses**

| Code | Description |
| ---- | ----------- |
| 200 |  |
| default |  |

##### ***PATCH***
**Summary:** Take a order, will update order status.

**Description:** Request:
default: takeRequest

**Parameters**

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| TakeRequest | body |  | Yes | [takeRequest](#takerequest) |

**Responses**

| Code | Description |
| ---- | ----------- |
| 200 |  |
| default |  |

### Models
---

### PlaceResponse  

This is used for returning a response with a single order as body

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| distance | long |  | No |
| id | long |  | No |
| status | string |  | No |

### placeRequest  

It is used to describe the initial order request.

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| destination | [ string ] | Destination latitude and longitude | Yes |
| origin | [ string ] | Origin latitude and longitude | Yes |

### takeRequest  

Status should be "TAKEN"

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| status | string | status should be "TAKEN" | Yes |