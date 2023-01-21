
CREATE TABLE vendors (
  ID STRING(36) NOT NULL,
  CreateTime TIMESTAMP NOT NULL,
  UpdatedTime TIMESTAMP NOT NULL
) PRIMARY KEY(ID);

CREATE TABLE variantGroups (
  ID STRING(36) NOT NULL,
  VendorID STRING(36) NOT NULL,
  CreateTime TIMESTAMP NOT NULL,
  UpdatedTime TIMESTAMP NOT NULL
)

INSERT INTO Products (ID, Title, Price, Stock, ContentOne, ProductVariantID, CreateTime, UpdateTime) values 
("29e6ad06-5fee-485e-894c-c5bbc73fa7d3", "商品タイトル", 10, 200, "内容内容内容", "7799d992-fe69-4812-8487-25f3e6d211ec", CURRENT_TIMESTAMP(), CURRENT_TIMESTAMP()); 