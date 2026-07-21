-- ============================================
-- CATEGORY 
-- ============================================
CREATE TABLE parsing_data.category (
    id BIGSERIAL PRIMARY KEY,
    marketplace VARCHAR(32) NOT NULL,
    name VARCHAR(255) NOT NULL,
    slug  TEXT NOT NULL,

    CONSTRAINT uq_marketplace_category
        UNIQUE (marketplace, slug)
);


-- =============================================
-- CATEGORY RELATION
-- ============================================
CREATE TABLE parsing_data.category_relation (
    parent_id BIGINT NOT NULL,
    child_id BIGINT NOT NULL,

    PRIMARY KEY(parent_id, child_id),

    FOREIGN KEY(parent_id)
        REFERENCES parsing_data.category(id),

    FOREIGN KEY(child_id)
        REFERENCES parsing_data.category(id)
);

-- ===========================================
--  QUERY
-- ===========================================
CREATE TABLE parsing_data.query (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    query VARCHAR(255) NOT NULL,

    CONSTRAINT uq_query_text
        UNIQUE (query)
);

-- ===========================================
--  BRAND 
-- ===========================================
CREATE TABLE parsing_data.brand (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    slug VARCHAR(128) NOT NULL,
    title VARCHAR(128) NOT NULL,

    CONSTRAINT uq_brand_slug
        UNIQUE (slug)
);


-- ================================================
-- SELLER 
-- ================================================
CREATE TABLE parsing_data.seller (
    id VARCHAR(20) PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    slug VARCHAR(100) NOT NULL,
    ogrn_ogrnip VARCHAR(50) NULL,
    inn VARCHAR(50) NOT NULL,

    CONSTRAINT uq_seller_id
        UNIQUE (id)
);


-- ================================================
-- BRAND - SELLER 
-- ================================================
CREATE TABLE parsing_data.brand_seller (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    seller_id VARCHAR(20) NOT NULL,
    brand_id INTEGER NOT NULL,

    CONSTRAINT uq_brand_seller
        UNIQUE (seller_id, brand_id),
    
    CONSTRAINT fk_brand_seller_brand_id
        FOREIGN KEY (brand_id)
        REFERENCES parsing_data.brand(id)
        ON DELETE CASCADE 
        ON UPDATE CASCADE,
    
    CONSTRAINT fk_brand_seller_seller_id
        FOREIGN KEY (seller_id)
        REFERENCES parsing_data.seller(id)
        ON DELETE CASCADE 
        ON UPDATE CASCADE
);


-- GOOD_ITEM 
-- =================================================
CREATE TABLE parsing_data.good_item (
    sku VARCHAR(50) PRIMARY KEY,
    slug VARCHAR(255) NOT NULL,
    title VARCHAR(255) NOT NULL,
    price INTEGER NULL,
    card_price INTEGER NULL,
    original_price INTEGER NULL,
    availability BOOLEAN NULL,
    seller_id VARCHAR(20) NOT NULL,
    brand_id INTEGER NULL,
    review_link VARCHAR(255) NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,

    CONSTRAINT uq_good_item_sku
        UNIQUE (sku),

    CONSTRAINT fk_good_seller
        FOREIGN KEY (seller_id)
        REFERENCES parsing_data.seller(id)
        ON DELETE SET NULL 
        ON UPDATE CASCADE,

    CONSTRAINT fk_good_brand
        FOREIGN KEY (brand_id)
        REFERENCES parsing_data.brand(id)
        ON DELETE SET NULL 
        ON UPDATE CASCADE
);

-- ===============================================
-- GOODS 
-- ===============================================
CREATE TABLE parsing_data.good (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    sku VARCHAR(50) NOT NULL,
    cat_id INTEGER NOT NULL,
    query_id INTEGER NOT NULL,
    glink VARCHAR(255) NOT NULL,

    CONSTRAINT uq_good_category
        UNIQUE (sku, cat_id, query_id),

    CONSTRAINT fk_good_category
        FOREIGN KEY (cat_id)
        REFERENCES parsing_data.category(id)
        ON DELETE RESTRICT
        ON UPDATE CASCADE,

    CONSTRAINT fk_good_query
        FOREIGN KEY (query_id)
        REFERENCES parsing_data.query(id)
        ON DELETE RESTRICT
        ON UPDATE CASCADE,

    CONSTRAINT fk_good_item
        FOREIGN KEY (sku)
        REFERENCES parsing_data.good_item(sku)
        ON DELETE CASCADE 
        ON UPDATE CASCADE
);


-- ===============================================
-- REVIEWS 
-- ===============================================
CREATE TABLE parsing_data.review (
    uuid UUID PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL,
    sku VARCHAR(50) NOT NULL,
    author_guid VARCHAR(50) NOT NULL,
    score INTEGER NOT NULL,
    "comment" TEXT NOT NULL,
    positive TEXT NOT NULL,
    negative TEXT NOT NULL,

    CONSTRAINT uq_review_uuid
        UNIQUE (uuid)
);



-- ===============================================
-- IMAGES 
-- ===============================================
CREATE TABLE parsing_data.image (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    sku VARCHAR(50) NOT NULL,
    img_url VARCHAR(255) NOT NULL,
    is_cover BOOLEAN NOT NULL,

    CONSTRAINT uq_image_url
        UNIQUE (img_url),

    CONSTRAINT fk_image_good_item
        FOREIGN KEY (sku)
        REFERENCES parsing_data.good_item(sku)
        ON DELETE CASCADE 
        ON UPDATE CASCADE
);


-- ===============================================
-- REVIEW IMAGES 
-- ===============================================
CREATE TABLE parsing_data.review_image (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    review_uuid UUID NOT NULL,
    url VARCHAR(255) NOT NULL,

    CONSTRAINT uq_review_image_url
        UNIQUE (url),

    CONSTRAINT fk_image_review
        FOREIGN KEY (review_uuid)
        REFERENCES parsing_data.review(uuid)
        ON DELETE CASCADE 
        ON UPDATE CASCADE
);
