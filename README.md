# Exoplanet Microservice

This Golang microservice is designed to manage information about different exoplanets, specifically Gas Giants and Terrestrial planets. It supports CRUD operations and fuel estimation for trips to these exoplanets.

## Features

The microservice provides the following functionalities:
1. **Add an Exoplanet**: Add a new exoplanet with properties like name, description, distance from Earth, radius, mass (only for Terrestrial planets), and type (Gas Giant or Terrestrial).
2. **List Exoplanets**: Retrieve a list of all available exoplanets.
3. **Get Exoplanet by ID**: Retrieve information about a specific exoplanet by its unique ID.
4. **Update Exoplanet**: Update the details of an existing exoplanet.
5. **Delete Exoplanet**: Remove an exoplanet from the catalog.
6. **Fuel Estimation**: Retrieve an overall fuel cost estimation for a trip to any particular exoplanet for a given crew capacity.

## Installation

1. **Clone the repository:**
    ```sh
    git clone <repository-url>
    cd <repository-name>
    ```

2. **Build the Docker image:**
    ```sh
    docker build -t exoplanet .
    ```

3. **Run the Docker container:**
    ```sh
    docker run -p 8080:8080 exoplanet
    ```

## Usage

Access the microservice at `http://localhost:8080`.

## API Endpoints

- **Add an Exoplanet**
  - **Endpoint:** `POST /exoplanets`
  - **Request Body:**
    ```json
    {
      "name": "string",
      "description": "string",
      "distance": "int",
      "radius": "double",
      "mass": "double (only for Terrestrial)",
      "type": "string (GasGiant or Terrestrial)"
    }
    ```
  - **Response:** `201 Created`

- **List Exoplanets**
  - **Endpoint:** `GET /exoplanets`
  - **Response:** `200 OK` with JSON array of exoplanets

- **Get Exoplanet by ID**
  - **Endpoint:** `GET /exoplanets/{id}`
  - **Response:** `200 OK` with JSON of the exoplanet

- **Update Exoplanet**
  - **Endpoint:** `PUT /exoplanets/{id}`
  - **Request Body:**
    ```json
    {
      "name": "string",
      "description": "string",
      "distance": "int",
      "radius": "double",
      "mass": "double (only for Terrestrial)",
      "type": "string (GasGiant or Terrestrial)"
    }
    ```
  - **Response:** `200 OK`

- **Delete Exoplanet**
  - **Endpoint:** `DELETE /exoplanets/{id}`
  - **Response:** `204 No Content`

- **Fuel Estimation**
  - **Endpoint:** `GET /exoplanets/{id}/fuel?crew={crew_capacity}`
  - **Response:** `200 OK` with fuel estimation
