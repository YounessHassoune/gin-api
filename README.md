# Gin Rest api 

This is a rest api project using Gin and Go created for learning purposes where we're going to try and implement different rest api concepts.

- [ ] Routing groups
- [ ] middlewares
- [ ] validation
- [ ] Auth
- [ ] validation 
- [ ] cors
- [ ] rate limit
- [ ] cache 
- [ ] i don't know ! many other cool stuff !

## Routes 

### User Routes:
- [ ] GET /products: Retrieve a list of all products.
- [ ] GET /products/{id}: Retrieve a specific product by its ID.
- [ ] POST /products: Create a new product.
- [ ] PUT /products/{id}: Update an existing product.
- [ ] DELETE /products/{id}: Delete a product.
### User Routes:
- [ ] POST /users/register: Register a new user.
- [ ] POST /users/login: Login a user.
- [ ] GET /users/profile: Retrieve user profile information.
- [ ] PUT /users/profile: Update user profile information.
### Cart Routes:
- [ ] GET /cart: Retrieve the items in the user's cart.
- [ ] POST /cart/add: Add a product to the user's cart.
- [ ] PUT /cart/update: Update the quantity of a product in the user's cart.
- [ ] DELETE /cart/remove/{id}: Remove a product from the user's cart.
### Order Routes:
- [ ] POST /orders: Place a new order.
- [ ] GET /orders: Retrieve a list of all orders.
- [ ] GET /orders/{id}: Retrieve a specific order by its ID.
- [ ] PUT /orders/{id}/cancel: Cancel an order.
- [ ] PUT /orders/{id}/update-status: Update the status of an order (e.g., processing, shipped, delivered).
### Payment Routes:
- [ ] POST /payment/charge: Process a payment for an order.
- [ ] POST /payment/webhook: Receive payment webhook events (e.g., payment success, failure).
### Category Routes:
- [ ] GET /categories: Retrieve a list of all product categories.
- [ ] GET /categories/{id}: Retrieve products under a specific category.
- [ ] POST /categories: Create a new product category.
- [ ] PUT /categories/{id}: Update an existing product category.
- [ ] DELETE /categories/{id}: Delete a product category.
### Review Routes:
- [ ] POST /products/{id}/reviews: Add a review for a product.
- [ ] GET /products/{id}/reviews: Retrieve reviews for a product.


