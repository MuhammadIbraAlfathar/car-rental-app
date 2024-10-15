# CAR RENTAL APP API

This is the Car Rental REST API documentation. This API is used to add customer data, update customer data, delete customers, add car data, update, and delete, then can make a booking for the car you want to rent.

## Table of Contents
- [Documentation](#documentation)
- [Tech Stack](#tech-stack)
- [Features v1](#features-v1)
- [Features v2](#features-v2)
- [Installation](#installation)


## Documentation

### Entity Relationship Diagram (ERD) v1
![image](https://github.com/MuhammadIbraAlfathar/assets/blob/main/erdv1.png?raw=true)


### Entity Relationship Diagram (ERD) v2
![image](https://github.com/MuhammadIbraAlfathar/assets/blob/main/erdv2.png?raw=true)



### Documentation API Testing PostmanAPI
- **https://documenter.getpostman.com/view/28873982/2sAXxTcW8H**


## Tech Stack
- **Golang**: The programming language used for the core backend logic.
- **Gin**: A high-performance HTTP web framework for building the RESTful API.
- **GORM**: An ORM library for Golang, used to interact with the PostgreSQL database.
- **PostgreSQL**: The primary database for storing all application data.
- **Docker**: Containerization for easy deployment and management.
- **Github Actions**: For build and deploy image docker.



## Features v1

- **Customer**: create, get customer by id, delete customer by id, get all customer.
- **Car**: create, get car by id, delete car by id, get all car.
- **Booking**: Make booking and calculate by start rent and end rent.


## Features v2

- **Customer**: create, get customer by id, delete customer by id, get all customer, there is a membership that is useful for discounts when booking.
- **Car**: create, get car by id, delete car by id, get all car.
- **Booking**: Make booking and calculate by start rent and end rent, calculate discount, calculate total driver cost, can add booking type and driver.



## Installation

Before you begin, ensure you have the these installed on your machine:
- Docker

### Steps
1. **Clone the Repository:**
    ```bash
    https://github.com/MuhammadIbraAlfathar/car-rental-app.git
   ```

2. **Navigate to the Project Directory:**
    ```bash
    cd car-rental-app
    ```

3. **Set Up Environment Variables:**  
   Create a `.env` file in the root directory and provide the necessary environment variables. See `.env.example` file for reference.

4. **Start the Server:**
   ```bash
   docker compose up -d --build
   ```