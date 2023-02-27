# Pochirfy is a EC platform that gathers new users and transfers user data to partner's EC site with UI specialized for first-time users.
Partners are only those vendors that have shopify EC site.
(this app don't actually sell products, and is not completed)

## words

- user: buyer
- customer (account): customer account of shopify
    - customer and email are one on one
        - (all order records are merged based on email within shopify)
    - activation
        - before activated
            - the case that a customer does not have a password
        - after activated
            - the case that a customer has a password
- payment vendor: paypay, fincode, applePay, googlePay servers
- partner: a vendor that sells actual products
- shopify admin: shopify admin api server

## Sequence

I. When user returns to pochirify-frontend after checkout
```mermaid
  sequenceDiagram
  participant pf as pochirify-frontend
  participant pb as pochirify-backend
  participant pv as payment vendor
  participant psa as partner shopify admin
  participant ps as partner shopify web
  autoNumber

  pf->>pb: call CreateOrder API
  pb->>psa: POST orders (create customer and pending order)
  psa-->>pb: 
  pb-->>pf: 
  pf->>pv: checkout
  pv-->>pf: 
  pf->>pb: call CompleteOrder API
  pb->>pv: check if checkout has been done
  pv-->>pb: 
  pb->>psa: call orderMarkAsPaid (update order as paid)
  psa-->>pb: 
  pb->>psa: POST account_activation_url (create shopify activation url)
  psa-->>pb: 
  pb-->>pf: return account_activation_url if customer not activated
  pf->>ps: transfer the user to a password setting page if the user consents
 
```

II. When user don’t return to pochirify-frontend after checkout
- (ex) when user kill pochirify-frontend after checkout with paypay
```mermaid
sequenceDiagram
participant pf as pochirify-frontend
participant pb as pochirify-backend
participant cs as cloud scheduler
participant pv as payment vendor
participant psa as partner shopify admin
participant ps as partner shopify
autoNumber

pf->>pb: call CreateOrder API
pb->>psa: POST orders(create customer and pending order)
psa-->>pb:  
pb-->>pf: 
pf->>pv: checkout
pv-->>pf: 
cs->>pb: call ReconcileOrders if checkout has been down
pb->>pv: check if checkout has been done
pv-->>pb: 
pb->psa: call orderMarkAsPaid
psa-->>pb:  
pb-->>cs: 

```

Ⅲ. when user don’t checkout after calling CreateOrder API
```mermaid
sequenceDiagram
participant pf as pochirify-frontend
participant pb as pochirify-backend
participant cs as cloud scheduler
participant pv as payment vendor
participant psa as partner shopify admin
participant ps as partner shopify
autoNumber

pf->>pb: call CreateOrder API
pb-->>pf: 
cs->>pb: call ReconcileOrders
pb->>pv: check if checkout has been done
pv-->>pb: 
pb->psa: POST cancel order
psa-->>pb: 
pb->>cs: 

```
