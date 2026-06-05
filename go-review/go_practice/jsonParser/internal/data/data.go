package data

var Data = `{
  "order_id": "ORD-2026-89431",
  "customer": {
    "id": "CUST-9921",
    "name": "Jane Doe",
    "is_premium_member": true,
    "contact": {
      "email": "jane.doe@example.com",
      "phone": "+1-555-0199"
    }
  },
  "shipping_address": {
    "street": "742 Evergreen Terrace",
    "city": "Springfield",
    "state": "IL",
    "postal_code": "62704",
    "coordinates": {
      "latitude": 39.7817,
      "longitude": -89.6501
    }
  },
  "items": [
    {
      "product_id": "PROD-001",
      "title": "Wireless Noise-Canceling Headphones",
      "quantity": 1,
      "price": 299.99,
      "tags": ["electronics", "audio", "wireless"],
      "discount": null
    },
    {
      "product_id": "PROD-042",
      "title": "Ergonomic Desk Chair",
      "quantity": 2,
      "price": 149.50,
      "tags": ["furniture", "office"],
      "discount": {
        "code": "WORKFROMHOME",
        "amount_saved": 30.00
      }
    }
  ],
  "payment": {
    "method": "Credit Card",
    "status": "Paid",
    "transaction_details": {
      "gateway": "Stripe",
      "card_brand": "Visa",
      "last_four": "4242"
    }
  }
}{
  "order_id": "ORD-2026-89431",
  "customer": {
    "id": "CUST-9921",
    "name": "Jane Doe",
    "is_premium_member": true,
    "contact": {
      "email": "jane.doe@example.com",
      "phone": "+1-555-0199"
    }
  },
  "shipping_address": {
    "street": "742 Evergreen Terrace",
    "city": "Springfield",
    "state": "IL",
    "postal_code": "62704",
    "coordinates": {
      "latitude": 39.7817,
      "longitude": -89.6501
    }
  },
  "items": [
    {
      "product_id": "PROD-001",
      "title": "Wireless Noise-Canceling Headphones",
      "quantity": 1,
      "price": 299.99,
      "tags": ["electronics", "audio", "wireless"],
      "discount": null
    },
    {
      "product_id": "PROD-042",
      "title": "Ergonomic Desk Chair",
      "quantity": 2,
      "price": 149.50,
      "tags": ["furniture", "office"],
      "discount": {
        "code": "WORKFROMHOME",
        "amount_saved": 30.00
      }
    }
  ],
  "payment": {
    "method": "Credit Card",
    "status": "Paid",
    "transaction_details": {
      "gateway": "Stripe",
      "card_brand": "Visa",
      "last_four": "4242"
    }
  }
}`