DROP TABLE IF EXISTS items;
CREATE TABLE items (
  id         INT AUTO_INCREMENT NOT NULL,
  name      VARCHAR(128) NOT NULL,
  price      DECIMAL(5,2) NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO items
  (name, price)
VALUES
  ('Fruit',56.99),
  ('Clothes', 63.99);