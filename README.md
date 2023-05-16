# Fiscally Fit

This is an API for calculating income tax based on salary and tax year. It provides a single endpoint that accepts two query parameters: income and year.

#API Specification

```yaml
openapi: 3.0.0
info:
  version: 1.0.0
  title: Tax Calculator API
  description: An API for calculating income tax based on salary and tax year
paths:
  /calculator:
    get:
      summary: Calculate income tax for a given salary and tax year
      parameters:
        - name: income
          in: query
          description: The annual income to calculate taxes for
          required: true
          schema:
            type: integer
            minimum: 0
        - name: year
          in: query
          description: The tax year to use for calculating taxes
          required: true
          schema:
            type: integer
            enum: [2019, 2020, 2021, 2022]
      responses:
        '200':
          description: Returns the total taxes owed for the salary
          content:
            application/json:
              schema:
                type: object
                properties:
                  total_taxes:
                    type: number
                  taxes_per_band:
                    type: object
                  effective_rate:
                    type: number
        '400':
          description: Invalid parameters provided
        '500':
          description: Internal server error
```

# Directory Structure
```

```
