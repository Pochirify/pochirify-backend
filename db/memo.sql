CREATE TABLE Users (
  ID STRING(36) NOT NULL,
  PhoneNumberDigest STRING(255) NOT NULL,
  IsAuthenticated BOOL NOT NULL,
  CreateTime TIMESTAMP NOT NULL,
  UpdateTime TIMESTAMP NOT NULL,
) PRIMARY KEY(ID);

CREATE TABLE UserAddresses (
  ID STRING(36) NOT NULL,
  UserID STRING(36) NOT NULL,
  EmailAddress STRING(319) NOT NULL,
  ZipCode INT64 NOT NULL,
  Prefecture STRING(4) NOT NULL,
  City STRING(255) NOT NULL,
  StreetAddress STRING(255) NOT NULL,
  Building STRING(255),
  LastName STRING(127) NOT NULL,
  FirstName STRING(127) NOT NULL,
  CreateTime TIMESTAMP NOT NULL,
  UpdateTime TIMESTAMP NOT NULL,
  CONSTRAINT FK_UserAddress_User_UserID FOREIGN KEY (UserID) REFERENCES Users (ID)
) PRIMARY KEY (ID);

CREATE TABLE Products (
  ID STRING(36) NOT NULL,
  Title STRING(255) NOT NULL,
  Price INT64 NOT NULL,
  Stock INT64 NOT NULL,
  ContentOne STRING(255) NOT NULL,
  ContentTwo STRING(255),
  ContentThree STRING(255),
  ContentFour STRING(255),
  ContentFive STRING(255),
  ProductVariantID String(36),
  CreateTime TIMESTAMP NOT NULL,
  UpdateTime TIMESTAMP NOT NULL
) PRIMARY KEY(ID);

CREATE TABLE Orders (
  ID STRING(36) NOT NULL,
  UserID STRING(36) NOT NULL,
  UserAddressID STRING(36) NOT NULL,
  Status STRING(10) NOT NULL,
  PaymentMethod STRING(20) NOT NULL,
  ProductID STRING(36) NOT NULL,
  Price INT64 NOT NULL,
  CreateTime TIMESTAMP NOT NULL,
  UpdateTime TIMESTAMP NOT NULL,
  CONSTRAINT FK_Order_User_UserID FOREIGN KEY (UserID) REFERENCES Users (ID),
  CONSTRAINT FK_Order_UserAddress_UserAddressID FOREIGN KEY (UserAddressID) REFERENCES UserAddresses (ID),
  CONSTRAINT FK_Order_Product_ProductID FOREIGN KEY (ProductID) REFERENCES Products (ID)
) PRIMARY KEY (ID);

# ------------------------------------------------
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