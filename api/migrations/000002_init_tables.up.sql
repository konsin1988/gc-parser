-- ============================================
-- CATEGORY 
-- ============================================
CREATE TABLE parsing_data.category (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE
);

-- ================================================
-- SELLER 
-- ================================================
CREATE TABLE parsing_data.seller (
    id INTEGER PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    ogrnip VARCHAR(255) NULL,
    inn VARCHAR(255) NOT NULL,

    CONSTRAINT uq_seller_id
        UNIQUE (id),
    CONSTRAINT uq_seller_inn
        UNIQUE (inn)
);


-- ===========================================
--  QUERY
-- ===========================================
CREATE TABLE parsing_data.query (
    id INTEGER PRIMARY KEY,
    cat_id INTEGER NOT NULL,
    query_text VARCHAR(255) NOT NULL,

    CONSTRAINT uq_query_category_text
        UNIQUE (cat_id, query_text),

    CONSTRAINT fk_query_category
        FOREIGN KEY (cat_id)
        REFERENCES parsing_data.category(id)
        ON DELETE RESTRICT
        ON UPDATE CASCADE
);

-- ==================================================
-- GOOD_ITEM 
-- =================================================
CREATE TABLE parsing_data.good_item (
    id INTEGER PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    price INTEGER NOT NULL,
    original_price INTEGER NULL,
    availability VARCHAR(100) NULL,
    seller_id INTEGER NOT NULL,

    CONSTRAINT uq_good_item_id
        UNIQUE (id),

    CONSTRAINT fk_good_seller
        FOREIGN KEY (seller_id)
        REFERENCES parsing_data.seller(id)
        ON DELETE RESTRICT
        ON UPDATE CASCADE
);

-- ===============================================
-- GOODS 
-- ===============================================
CREATE TABLE parsing_data.good (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    good_id INTEGER NOT NULL,
    cat_id INTEGER NOT NULL,
    glink VARCHAR(255) NOT NULL,

    CONSTRAINT uq_good_category
        UNIQUE (good_id, cat_id),

    CONSTRAINT fk_good_category
        FOREIGN KEY (cat_id)
        REFERENCES parsing_data.category(id)
        ON DELETE RESTRICT
        ON UPDATE CASCADE
);


-- ===============================================
-- COMMENTS 
-- ===============================================
CREATE TABLE parsing_data."comment" (
    uuid UUID PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL,
    good_id INTEGER NOT NULL,
    author_guid VARCHAR(50) NOT NULL,
    "comment" VARCHAR(255) NOT NULL,
    positive VARCHAR(255) NOT NULL,
    negative VARCHAR(255) NOT NULL,
    ph_urls VARCHAR(255) NULL,

    CONSTRAINT uq_comment_uuid
        UNIQUE (uuid),

    CONSTRAINT fk_comment_good
        FOREIGN KEY (good_id)
        REFERENCES parsing_data.good(id)
        ON DELETE RESTRICT
        ON UPDATE CASCADE
);
