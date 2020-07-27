# IMS-BE

Backend for Inventory Management System

**Cài đặt GORM**
> go get -u github.com/jinzhu/gorm

**Biến môi trường**
* API_HOST_PORT = 0.0.0.0:5000
* DB_HOST = postgres
* DB_PORT = 5432
* DB_NAME = inventory
* DB_USER = admin
* DB_PASSWORD = inventory

**API**
> Kiểu gửi của phương thức POST, PUT là form-data

* GET /api/products(?page={page})
    * nếu không có phần trong ngoặc, trả về tất cả sản phẩm
    * nếu có trả về page, mỗi page 5 sản phẩm
* POST /api/products
    * Request Headers: 
    > Content-Type : application/json
    * JSON:
    > {\
         "prod_name":"text_here"\
        }
* GET /api/products/{id}
* PUT /api/products/{id}
    * Request Headers: 
    > Content-Type : application/json
    * JSON:
    > {\
         "prod_name":"text_here"\
        }
* DELETE /api/products/{id}
* GET /api/products/search?name={text}&(page={page})
* GET /api/imports(?page={page})
* POST /api/imports
    * Request Headers: 
        > Content-Type : application/json
        * JSON:
        > {\
             "prod_id": {integer_here},\
             "price": {float_here},\
             "quantity": {integer_here}\
            }
* GET /api/imports/{id}
* GET /api/exports(?page={page})
* POST /api/exports
    * Request Headers: 
    > Content-Type : application/json
    * JSON:
    > {\
         "prod_id": {integer_here},\
         "quantity": {integer_here}\
        }
* GET /api/exports/{id}
* GET /api/inventories(?page={page})
* GET /api/inventories/{id}
* GET /api/inventories/search?name={text}&(page={page})

    
**Code**
* 404 : Not Found
* 405 : Method not support
* 606 : Cannot delete product because there is still product in Inventory
* 666 : Null Value
* 777 : Existed

