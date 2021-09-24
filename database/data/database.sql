-- MySQL Script generated by MySQL Workbench
-- Fri Sep 24 13:52:28 2021
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema evermos-online
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema evermos-online
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `evermos-online` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci ;
USE `evermos-online` ;

-- -----------------------------------------------------
-- Table `evermos-online`.`suppliers`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `evermos-online`.`suppliers` (
  `supplier_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255) NOT NULL,
  `email` VARCHAR(100) NOT NULL,
  `password` VARCHAR(255) NOT NULL,
  `last_login_at` TIMESTAMP(6) NULL DEFAULT NULL,
  `is_active` CHAR(1) NOT NULL DEFAULT '1',
  `created_at` TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
  `updated_at` TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
  `deleted_at` TIMESTAMP(6) NULL DEFAULT NULL,
  PRIMARY KEY (`supplier_id`),
  UNIQUE INDEX `email` (`email` ASC) VISIBLE,
  INDEX `idx_supp_status` (`is_active` ASC) VISIBLE,
  INDEX `idx_supp_name` (`name` ASC) VISIBLE)
ENGINE = InnoDB
AUTO_INCREMENT = 1
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `evermos-online`.`products`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `evermos-online`.`products` (
  `inv_id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `supplier_id` INT UNSIGNED NOT NULL,
  `name` VARCHAR(255) NOT NULL,
  `description` VARCHAR(255) NOT NULL,
  `qty_stock` INT UNSIGNED NOT NULL DEFAULT '0',
  `qty_min_stock` INT UNSIGNED NOT NULL DEFAULT '1',
  `qty_min_cart` INT UNSIGNED NOT NULL DEFAULT '1',
  `qty_pavorit` INT UNSIGNED NOT NULL DEFAULT '0',
  `qty_like` INT UNSIGNED NOT NULL DEFAULT '0',
  `amount_price` INT UNSIGNED NOT NULL DEFAULT '0',
  `amount_disc` INT UNSIGNED NOT NULL DEFAULT '0',
  `is_promo` CHAR(1) NOT NULL DEFAULT '0',
  `is_active` CHAR(1) NOT NULL DEFAULT '1',
  `created_at` TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
  `updated_at` TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
  `deleted_at` TIMESTAMP(6) NULL DEFAULT NULL,
  PRIMARY KEY (`inv_id`),
  UNIQUE INDEX `name` (`name` ASC) VISIBLE,
  INDEX `idx_inv_qty_stock` (`qty_stock` ASC) VISIBLE,
  INDEX `idx_inv_status` (`is_active` ASC) VISIBLE,
  INDEX `fk_products_suppliers_idx` (`supplier_id` ASC) VISIBLE,
  INDEX `idx_inv_promo` (`is_promo` ASC) VISIBLE,
  INDEX `idx_product_deleted_at` (`deleted_at` ASC) VISIBLE,
  CONSTRAINT `fk_products_suppliers`
    FOREIGN KEY (`supplier_id`)
    REFERENCES `evermos-online`.`suppliers` (`supplier_id`)
    ON DELETE RESTRICT
    ON UPDATE RESTRICT)
ENGINE = InnoDB
AUTO_INCREMENT = 1
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `evermos-online`.`users`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `evermos-online`.`users` (
  `user_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255) NOT NULL,
  `email` VARCHAR(100) NOT NULL,
  `password` VARCHAR(255) NOT NULL,
  `last_login_at` TIMESTAMP(6) NULL DEFAULT NULL,
  `is_active` CHAR(1) NOT NULL DEFAULT '1',
  `created_at` TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
  `updated_at` TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
  `deleted_at` TIMESTAMP(6) NULL DEFAULT NULL,
  PRIMARY KEY (`user_id`),
  UNIQUE INDEX `email` (`email` ASC) VISIBLE,
  INDEX `idx_user_status` (`is_active` ASC) VISIBLE,
  INDEX `idx_user_deleted_at` (`deleted_at` ASC) VISIBLE)
ENGINE = InnoDB
AUTO_INCREMENT = 1
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `evermos-online`.`carts`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `evermos-online`.`carts` (
  `cart_id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` INT UNSIGNED NOT NULL,
  `session_id` VARCHAR(45) NULL,
  `qty_item` INT(10) UNSIGNED NOT NULL DEFAULT '0',
  `amount_price` INT(10) UNSIGNED NOT NULL DEFAULT '0',
  `amount_disc` INT(10) UNSIGNED NOT NULL DEFAULT '0',
  `amount_total` INT UNSIGNED NOT NULL DEFAULT '0',
  `checkout_date` TIMESTAMP(6) NULL,
  `is_active` CHAR(1) NOT NULL DEFAULT '1',
  `created_at` TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
  `updated_at` TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
  `deleted_at` TIMESTAMP(6) NULL DEFAULT NULL,
  PRIMARY KEY (`cart_id`),
  INDEX `idx_cartsr_status` (`is_active` ASC) VISIBLE,
  INDEX `fk_carts_user_idx` (`user_id` ASC) VISIBLE,
  INDEX `idx_carts_session_id` (`session_id` ASC) INVISIBLE,
  INDEX `idx_carts_deleted_at` (`deleted_at` ASC) VISIBLE,
  CONSTRAINT `fk_carts_users`
    FOREIGN KEY (`user_id`)
    REFERENCES `evermos-online`.`users` (`user_id`))
ENGINE = InnoDB
AUTO_INCREMENT = 1
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `evermos-online`.`cart_details`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `evermos-online`.`cart_details` (
  `cart_id` BIGINT(19) UNSIGNED NOT NULL,
  `cart_no` SMALLINT(2) UNSIGNED NOT NULL DEFAULT '1',
  `inv_id` BIGINT UNSIGNED NOT NULL,
  `supplier_id` INT UNSIGNED NOT NULL,
  `name` VARCHAR(255) NOT NULL DEFAULT 'product item name',
  `noted` VARCHAR(255) NULL,
  `qty_order` INT UNSIGNED NOT NULL DEFAULT '0',
  `amount_price` INT UNSIGNED NOT NULL DEFAULT '0',
  `amount_disc` INT UNSIGNED NOT NULL DEFAULT '0',
  `amount_total` INT NOT NULL DEFAULT '1',
  `is_promo` CHAR(1) NOT NULL DEFAULT '0',
  `is_active` CHAR(1) NOT NULL DEFAULT '1',
  `created_at` TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
  `updated_at` TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
  `deleted_at` TIMESTAMP(6) NULL DEFAULT NULL,
  PRIMARY KEY (`cart_id`, `cart_no`),
  INDEX `idx_inv_status` (`is_active` ASC) VISIBLE,
  INDEX `fk_cart_items_carts_idx` (`cart_id` ASC) VISIBLE,
  INDEX `fk_cart_items_products_idx` (`inv_id` ASC) VISIBLE,
  INDEX `fk_cart_items_suppliers_idx` (`supplier_id` ASC) VISIBLE,
  INDEX `idx_cart_items_deleted_at` (`deleted_at` ASC) VISIBLE,
  CONSTRAINT `fk_cart_details_carts`
    FOREIGN KEY (`cart_id`)
    REFERENCES `evermos-online`.`carts` (`cart_id`)
    ON DELETE RESTRICT
    ON UPDATE RESTRICT,
  CONSTRAINT `fk_cart_details_products`
    FOREIGN KEY (`inv_id`)
    REFERENCES `evermos-online`.`products` (`inv_id`)
    ON DELETE RESTRICT
    ON UPDATE RESTRICT,
  CONSTRAINT `fk_cart_details_suppliers`
    FOREIGN KEY (`supplier_id`)
    REFERENCES `evermos-online`.`suppliers` (`supplier_id`)
    ON DELETE RESTRICT
    ON UPDATE RESTRICT)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `evermos-online`.`status_order`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `evermos-online`.`status_order` (
  `status_order` CHAR(1) NOT NULL,
  `name` VARCHAR(255) NOT NULL,
  `is_active` CHAR(1) NOT NULL DEFAULT '1',
  `created_at` TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
  `updated_at` TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
  `deleted_at` TIMESTAMP(6) NULL DEFAULT NULL,
  PRIMARY KEY (`status_order`),
  INDEX `idx_status_order_status` (`is_active` ASC) VISIBLE,
  UNIQUE INDEX `name_UNIQUE` (`name` ASC) VISIBLE)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `evermos-online`.`sales`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `evermos-online`.`sales` (
  `sales_id` BIGINT(19) UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` INT UNSIGNED NOT NULL,
  `user_address` VARCHAR(255) NOT NULL,
  `sales_inv_no` VARCHAR(45) NOT NULL,
  `sales_date` TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
  `promo_code` VARCHAR(25) NULL,
  `qty_item` INT(10) UNSIGNED NOT NULL DEFAULT '0',
  `amount_price` INT(10) UNSIGNED NOT NULL DEFAULT '0',
  `amount_disc` INT(10) UNSIGNED NOT NULL DEFAULT '0',
  `amount_expedition` INT UNSIGNED NOT NULL DEFAULT '0',
  `amount_token` INT UNSIGNED NOT NULL DEFAULT '0',
  `amount_total` INT UNSIGNED NOT NULL DEFAULT '0',
  `amount_payment` INT UNSIGNED NOT NULL DEFAULT '0',
  `payment_date` TIMESTAMP(6) NULL,
  `status_order` CHAR(1) NOT NULL,
  `is_active` CHAR(1) NOT NULL DEFAULT '1',
  `created_at` TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
  `updated_at` TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
  `deleted_at` TIMESTAMP(6) NULL DEFAULT NULL,
  PRIMARY KEY (`sales_id`),
  INDEX `idx_sales_status` (`is_active` ASC) VISIBLE,
  INDEX `fk_sales_users_idx` (`user_id` ASC) VISIBLE,
  INDEX `idx_sales_order_date` (`sales_date` ASC) VISIBLE,
  INDEX `idx_sales_deleted_at` (`deleted_at` ASC) VISIBLE,
  INDEX `fk_sales_status_order_idx` (`status_order` ASC) VISIBLE,
  CONSTRAINT `fk_sales_users`
    FOREIGN KEY (`user_id`)
    REFERENCES `evermos-online`.`users` (`user_id`),
  CONSTRAINT `fk_sales_status_order`
    FOREIGN KEY (`status_order`)
    REFERENCES `evermos-online`.`status_order` (`status_order`)
    ON DELETE RESTRICT
    ON UPDATE RESTRICT)
ENGINE = InnoDB
AUTO_INCREMENT = 1
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `evermos-online`.`sales_details`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `evermos-online`.`sales_details` (
  `sales_id` BIGINT(19) UNSIGNED NOT NULL,
  `sales_no` SMALLINT(2) UNSIGNED NOT NULL DEFAULT '1',
  `inv_id` BIGINT UNSIGNED NOT NULL,
  `supplier_id` INT UNSIGNED NOT NULL,
  `name` VARCHAR(255) NOT NULL DEFAULT 'product item name',
  `qty_order` INT UNSIGNED NOT NULL DEFAULT '0',
  `amount_price` INT UNSIGNED NOT NULL DEFAULT '0',
  `amount_disc` INT UNSIGNED NOT NULL DEFAULT '0',
  `amount_total` INT UNSIGNED NOT NULL DEFAULT '1',
  `noted` VARCHAR(255) NULL,
  `is_promo` CHAR(1) NOT NULL DEFAULT '0',
  `is_active` CHAR(1) NOT NULL DEFAULT '1',
  `created_at` TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
  `updated_at` TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
  `deleted_at` TIMESTAMP(6) NULL DEFAULT NULL,
  PRIMARY KEY (`sales_id`, `sales_no`),
  INDEX `idx_sales_items_status` (`is_active` ASC) INVISIBLE,
  INDEX `fk_sales_items_products_idx` (`inv_id` ASC) VISIBLE,
  INDEX `fk_sales_items_suppliers_idx` (`supplier_id` ASC) VISIBLE,
  CONSTRAINT `fk_sales_details_sales`
    FOREIGN KEY (`sales_id`)
    REFERENCES `evermos-online`.`sales` (`sales_id`)
    ON DELETE RESTRICT
    ON UPDATE RESTRICT,
  CONSTRAINT `fk_sales_details_products`
    FOREIGN KEY (`inv_id`)
    REFERENCES `evermos-online`.`products` (`inv_id`)
    ON DELETE RESTRICT
    ON UPDATE RESTRICT,
  CONSTRAINT `fk_sales_details_suppliers`
    FOREIGN KEY (`supplier_id`)
    REFERENCES `evermos-online`.`suppliers` (`supplier_id`)
    ON DELETE RESTRICT
    ON UPDATE RESTRICT)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `evermos-online`.`order_types`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `evermos-online`.`order_types` (
  `order_type` CHAR(1) NOT NULL,
  `name` VARCHAR(255) NOT NULL,
  `is_active` CHAR(1) NOT NULL DEFAULT '1',
  `created_at` TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
  `updated_at` TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
  `deleted_at` TIMESTAMP(6) NULL DEFAULT NULL,
  PRIMARY KEY (`order_type`),
  INDEX `idx_order_types_status` (`is_active` ASC) VISIBLE,
  UNIQUE INDEX `name_UNIQUE` (`name` ASC) VISIBLE)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `evermos-online`.`product_orders`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `evermos-online`.`product_orders` (
  `po_id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `inv_id` BIGINT UNSIGNED NOT NULL,
  `supplier_id` INT UNSIGNED NOT NULL,
  `order_type` CHAR(1) NOT NULL,
  `order_date` TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
  `qty_order` INT UNSIGNED NOT NULL DEFAULT '0',
  `amount_price` INT UNSIGNED NOT NULL DEFAULT '0' COMMENT 'price per item',
  `amount_disc` INT UNSIGNED NOT NULL DEFAULT '0' COMMENT 'discount per item',
  `is_promo` CHAR(1) NOT NULL DEFAULT '0',
  INDEX `idx_products_order_date` (`order_date` ASC) VISIBLE,
  INDEX `fk_products_orders_products_idx` (`inv_id` ASC) VISIBLE,
  PRIMARY KEY (`po_id`),
  INDEX `fk_products_orders_suppliers_idx` (`supplier_id` ASC) VISIBLE,
  INDEX `fk_products_orders_order_types1_idx` (`order_type` ASC) VISIBLE,
  CONSTRAINT `fk_products_orders_products1`
    FOREIGN KEY (`inv_id`)
    REFERENCES `evermos-online`.`products` (`inv_id`)
    ON DELETE RESTRICT
    ON UPDATE RESTRICT,
  CONSTRAINT `fk_products_orders_suppliers1`
    FOREIGN KEY (`supplier_id`)
    REFERENCES `evermos-online`.`suppliers` (`supplier_id`)
    ON DELETE RESTRICT
    ON UPDATE RESTRICT,
  CONSTRAINT `fk_products_orders_order_types1`
    FOREIGN KEY (`order_type`)
    REFERENCES `evermos-online`.`order_types` (`order_type`)
    ON DELETE RESTRICT
    ON UPDATE RESTRICT)
ENGINE = InnoDB
AUTO_INCREMENT = 1
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;

USE `evermos-online`;

DELIMITER $$
USE `evermos-online`$$
CREATE DEFINER = CURRENT_USER TRIGGER `evermos-online`.`sales_details_AFTER_INSERT` AFTER INSERT ON `sales_details` FOR EACH ROW
BEGIN
	if New.qty_order > 0 then
		update `evermos-online`.`products` set qty_stock = qty_stock-New.qty_order where inv_id= New.inv_id and supplier_id=New.supplier_id;
	end if;
END$$


DELIMITER ;

SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;