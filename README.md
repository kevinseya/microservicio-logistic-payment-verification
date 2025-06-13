# microservicio-logistic-payment-verification
# Go Project: Logistics Payment Verification

## microservice-logistic-payment-verification

This is a microservice developed in **Go** that provides an API to verify payments in a logistics system. Its main functionality is to validate payments through Stripe and report the result via a webhook.

## Prerequisites

Make sure you have the following installed on your system:

- **Go** (v1.18 or higher)
- **Docker**
- **MySQL** (if data storage is required)
- **Stripe** account and credentials

## Setup

### 1. Clone the repository

If the project is hosted on a Git repository, clone it to your local machine:

```sh
git clone https://github.com/kevinseya/microservicio-logistic-payment-verification.git
```

### 2. Configure environment variables

Create a `.env` file in the root of the project with the following variables:

```sh
STRIPE_SECRET_KEY=your_stripe_secret_key
NOTIFICATION_WEBHOOK_PAYMENT_URL=your_webhook_url
```

### 3. Running the project

To run the server locally, use the following command:

```sh
go run main.go
```

The server will start on port **8088**. You can test it with:

```
http://localhost:8088/api/payment/validation
```

## Project Structure

```
.microservice-logistic-payment-verification/
├── .github/workflows/dockerhub_ec2.yml # CI/CD Configuration
├── config/
│ ├── config.go # Loading Environment Variables
├── controller/
│ ├── payment_controller.go # Payment Controller
├── model/
│ ├── payment.go # Payment Data Model
├── routes/
│ ├── routes.go # Route Definition
├── service/
│ ├── payment_service.go # Stripe payment validation logic
├── .gitignore
├── Dockerfile # Docker configuration
├── go.mod # Project dependencies
├── go.sum # Dependency checksum
├── main.go # Application entry point
└── README.md
```

## CI/CD with Docker and EC2

This microservice is configured to be deployed on **AWS EC2** using GitHub Actions.

### 1. Building and Pushing Image to Docker Hub

The `.github/workflows/dockerhub_ec2.yml` file contains the steps to:
- Authenticate to Docker Hub
- Build the image with `docker build`
- Push it to Docker Hub with `docker push`

### 2. Deploy to EC2

- The EC2 instance fetches the image from Docker Hub
- Any previous containers are stopped and removed
- A `.env` file with the necessary credentials is generated
- The container is started with `docker run`

## Docker

To build and run the microservice in Docker manually:

### 1. Build the image
```sh
docker build -t logistic_payment_verification .
```

### 2. Run the container
```sh
docker run -d --name logistic_payment_verification -p 8088:8088 --env-file .env logistic_payment_verification
```

## API Endpoints

### Validate Payment
- **Route:** `POST /api/payment/validation`
- **Body:**

```json
{
"payment_intent": "pi_1234567890"
}
```

- **Response:**
```json
{
"valid": true
}
```

## Technologies Used

- **Go** (Golang) as the main language
- **Stripe API** for payment processing
- **Docker** for containerization
- **GitHub Actions** for CI/CD
- **AWS EC2** for deployment in the cloud

