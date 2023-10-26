CREATE TABLE IF NOT EXISTS client
(
    "id"       serial NOT NULL,
    "name"     varchar(150) NOT NULL,
    cpf        varchar(30) NULL,
    email      varchar(50) NULL,
    created_at timestamp DEFAULT NOW(),
    CONSTRAINT PK_client PRIMARY KEY ( "id" )
);

CREATE TABLE IF NOT EXISTS product
(
    "id"           serial NOT NULL,
    "name"         varchar(150) NOT NULL,
    category       varchar(100) NOT NULL,
    image          varchar(255) NOT NULL,
    description    text NOT NULL,
    price          decimal(8,2) NULL,
    created_at     timestamp DEFAULT NOW(),
    deactivated_at timestamp DEFAULT NULL,
    CONSTRAINT PK_product PRIMARY KEY ( "id" )
);

CREATE TABLE IF NOT EXISTS voucher
(
    "id"         serial NOT NULL,
    code       varchar(150) NOT NULL,
    percentage integer NOT NULL,
    created_at timestamp DEFAULT NOW(),
    expires_at timestamp NOT NULL,
    CONSTRAINT PK_voucher PRIMARY KEY ( "id" )
);

CREATE TABLE IF NOT EXISTS orders
(
    "id"                serial NOT NULL,
    status            varchar(30) NOT NULL,
    verification_code varchar(50) NOT NULL,
    created_at        timestamp DEFAULT NOW(),
    client_id         int NOT NULL,
    voucher_id        int NULL,
    CONSTRAINT PK_orders PRIMARY KEY ( "id" ),
    CONSTRAINT FK_client_id FOREIGN KEY ( client_id ) REFERENCES client ( "id" ),
    CONSTRAINT FK_voucher_id FOREIGN KEY ( voucher_id ) REFERENCES voucher ( "id" )
);
CREATE INDEX FK_client_id ON orders (client_id);
CREATE INDEX FK_voucher_id ON orders (voucher_id);


CREATE TABLE IF NOT EXISTS orders_products
(
    "id"          serial NOT NULL,
    quantity    int NOT NULL,
    total_price decimal(8,2) NULL,
    discount    decimal(8,2) DEFAULT 0,
    orders_id   int NOT NULL,
    product_id  int NOT NULL,
    created_at timestamp DEFAULT NOW(),
    CONSTRAINT PK_orders_products PRIMARY KEY ( "id" ),
    CONSTRAINT FK_orders_id FOREIGN KEY ( orders_id ) REFERENCES orders ( "id" ),
    CONSTRAINT FK_product_id FOREIGN KEY ( product_id ) REFERENCES product ( "id" )
);
CREATE INDEX FK_orders_id ON orders_products (orders_id);
CREATE INDEX FK_product_id ON orders_products (product_id);
