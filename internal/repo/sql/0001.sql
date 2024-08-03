CREATE TABLE IF NOT EXISTS `contract` (
	`id` BIGINT NOT NULL,
	`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`address` VARCHAR(66) NOT NULL,
	`schema` VARCHAR(32) NOT NULL,
	`name` VARCHAR(255) NOT NULL,
	`created_at_block_number` BIGINT UNSIGNED NOT NULL,
	`synced_at_block_number` BIGINT UNSIGNED NOT NULL,
	`processed_at_block_number` BIGINT UNSIGNED NOT NULL,
	`processed_at_log_index` BIGINT NOT NULL,
	PRIMARY KEY (`id`),
	CONSTRAINT uk_address UNIQUE (`address`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `ethereum_log` (
	`address` VARCHAR(66) NOT NULL,
	`block_number` BIGINT UNSIGNED NOT NULL,
	`log_index` INT UNSIGNED NOT NULL,
	`tx_hash` VARCHAR(66) NOT NULL,
	`tx_index` INT UNSIGNED NOT NULL,
	`block_hash` VARCHAR(66) NOT NULL,
	`removed` TINYINT(1) NOT NULL,
	`topics` JSON NOT NULL,
	`data` TEXT NOT NULL,
	PRIMARY KEY (`block_number`, `log_index`),
	INDEX idx_address (`address`, `block_number`, `log_index`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `nft_item` (
	`id` BIGINT(20) NOT NULL,
	`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`collection_address` CHAR(66) NOT NULL,
	`owner_address` CHAR(66) NOT NULL,
	`token_index` BIGINT(20) NOT NULL,
	`token_id` VARCHAR(66) NOT NULL,
    `meta_data` TEXT NULL,
	`token_url` TEXT NULL,
    `name` VARCHAR(255) NOT NULL,
    `description` TEXT NULL,
    `nft_price` VARCHAR(255) NOT NULL,
	PRIMARY KEY (id),
	INDEX idx_nft_item (collection_address,token_id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

CREATE TABLE IF NOT EXISTS `order` (
    `id` BIGINT(20) NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `collection_address` CHAR(66) NOT NULL,
    `seller_address` CHAR(66) NOT NULL,
    `buyer_address` CHAR(66) NOT NULL,
    `token_id` VARCHAR(66) NOT NULL,
    `token_url` TEXT NULL,
    `price` VARCHAR(255) NOT NULL,
    `tx_hash` VARCHAR(66) NOT NULL,
    PRIMARY KEY (id),
    INDEX idx_nft_item (collection_address,token_id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

CREATE TABLE IF NOT EXISTS `market` (
    `id` BIGINT(20) NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `collection_address` CHAR(66) NOT NULL,
    `seller_address` CHAR(66) NOT NULL,
    `token_id` VARCHAR(66) NOT NULL,
    `token_url` TEXT NULL,
    `price` VARCHAR(255) NOT NULL,
    `price_float` DECIMAL(20, 18) NOT NULL,
    `status` TINYINT(1) NOT NULL,
    PRIMARY KEY (id),
    INDEX idx_nft_item (collection_address,token_id),
    INDEX idx_nft_item_seller (collection_address,token_id,seller_address),
    INDEX idx_nft_item_status (collection_address,token_id,seller_address,status)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

CREATE TABLE IF NOT EXISTS `ethereum_block_header` (
	`hash` VARCHAR(66) NOT NULL,
	`parent_hash` VARCHAR(66) NOT NULL,
	`sha3_uncles` VARCHAR(66) NOT NULL,
	`miner` VARCHAR(42) NOT NULL,
	`state_root` VARCHAR(66) NOT NULL,
	`transactions_root` VARCHAR(66) NOT NULL,
	`receipts_root` VARCHAR(66) NOT NULL,
	`logs_bloom` TEXT NOT NULL,
	`difficulty` BIGINT NULL,
	`number` BIGINT NULL,
	`gas_limit` BIGINT NOT NULL,
	`gas_used` BIGINT NOT NULL,
	`timestamp` BIGINT NOT NULL,
	`extra_data` TEXT NOT NULL,
	`mix_hash` VARCHAR(66) NOT NULL,
	`nonce` BIGINT NOT NULL,
	`base_fee_per_gas` BIGINT NULL,
	`withdrawals_root` VARCHAR(66) NULL,
	PRIMARY KEY (`hash`),
	INDEX idx_number (`number`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;