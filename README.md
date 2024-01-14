# RestQL
[![CodeQL](https://github.com/gogean/restql/actions/workflows/codeql.yml/badge.svg)](https://github.com/gogean/restql/actions/workflows/codeql.yml)
RestQL is a library that automatically generates REST APIs based on the provided configuration. Think of this library's API as a configuration itself.

## Configuration

1. **Data Source**

   - Database Connection String
   - Database Query Template
      - Template name
      - Input schema for validation

2. **API Schema** 

   A JSON Schema designed to specify the properties of a generated API. It encompasses key attributes, including:
   - API path
   - HTTP method
   - Input schema for defining and validating the structure of the API payload
   - Datasource template selection





