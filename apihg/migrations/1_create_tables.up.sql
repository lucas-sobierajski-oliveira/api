CREATE TABLE products (
	id int AUTO_INCREMENT PRIMARY KEY,
	name varchar(50),
    product varchar(50) NOT NULL
);

CREATE TABLE plans (
	id int(5) AUTO_INCREMENT PRIMARY KEY,
    price_renew float NOT NULL,
    price_order float NOT NULL,
    month int NOT NULL,
    description varchar(50) NOT NULL,
    id_product integer not null,
    CONSTRAINT fk_product_id FOREIGN KEY (id_product) references products (id)
);