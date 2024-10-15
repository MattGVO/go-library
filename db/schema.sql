DROP TABLE IF EXISTS "public"."books";

CREATE TABLE "public"."books" (
    "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
    "title" varchar(255) NOT NULL,
    "author" varchar(255) NOT NULL,
    "genre" varchar(100),
    "published_year" int4,
    PRIMARY KEY ("id")
);

DROP TABLE IF EXISTS "public"."checkouts";

CREATE TABLE "public"."checkouts" (
    "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
    "copy_id" uuid NOT NULL,
    "user_id" uuid NOT NULL,
    "checkout_date" date NOT NULL,
    "return_date" date,
    "due_date" date NOT NULL,
    PRIMARY KEY ("id")
);

DROP TABLE IF EXISTS "public"."copies";

CREATE TABLE "public"."copies" (
    "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
    "book_id" uuid NOT NULL,
    "acquired_date" date DEFAULT CURRENT_DATE,
    "edition" varchar(50),
    "condition" varchar(50),
    PRIMARY KEY ("id")
);

DROP TABLE IF EXISTS "public"."users";

CREATE TABLE "public"."users" (
    "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
    "full_name" varchar(255) NOT NULL,
    "email" varchar(255) NOT NULL,
    "phone_number" varchar(20),
    "registered_date" date DEFAULT CURRENT_DATE,
    PRIMARY KEY ("id")
);

INSERT INTO "public"."books" ("id", "title", "author", "genre", "published_year") VALUES
('10953403-56f8-4c6b-9926-b7f43a2bb35a', 'The Hobbit', 'J.R.R. Tolkien', 'Fantasy', 1937),
('1aac4b04-035e-40b6-890c-6bdb735425ab', '1984', 'George Orwell', 'Dystopian', 1949),
('22aee4e5-a8f5-43da-8af4-6386ac990798', 'The Catcher in the Rye', 'J.D. Salinger', 'Fiction', 1951),
('26f85cce-b193-4da1-bc96-2e2c588604de', 'Pride and Prejudice', 'Jane Austen', 'Romance', 1813),
('4683079b-fb7d-410f-9ef2-1412185d4108', 'The Shining', 'Stephen King', 'Horror', 1977),
('52096833-d586-4846-ab4c-450eb9dda592', 'The Great Gatsby', 'F. Scott Fitzgerald', 'Fiction', 1925),
('7dd2ca88-4ecb-4d45-a874-7591ce455708', 'Crime and Punishment', 'Fyodor Dostoevsky', 'Psychological Fiction', 1866),
('91c54736-2882-422e-b322-5365722c1641', 'Harry Potter and the Sorcerer’s Stone', 'J.K. Rowling', 'Fantasy', 1997),
('9f0161e3-42d4-4122-892b-dd466ff54225', 'The Da Vinci Code', 'Dan Brown', 'Mystery', 2003),
('ba0a97ff-6cae-4ba9-b6a4-cb5b3e386c2d', 'The Alchemist', 'Paulo Coelho', 'Adventure', 1988),
('e1e0f7d9-d17f-454d-9446-6834c804ba46', 'Brave New World', 'Aldous Huxley', 'Dystopian', 1932),
('e5d5713d-35fb-4981-b7ae-ed1db45723a5', 'To Kill a Mockingbird', 'Harper Lee', 'Fiction', 1960),
('e9b0de0d-8fa3-4bb2-abd0-6be71f9a4040', 'The Lord of the Rings: The Fellowship of the Ring', 'J.R.R. Tolkien', 'Fantasy', 1954);

INSERT INTO "public"."checkouts" ("id", "copy_id", "user_id", "checkout_date", "return_date", "due_date") VALUES
('0d06d804-fa04-40ec-9713-1df5b63c6c66', 'db018be1-75b3-4ef3-b59c-dcd3fd44232c', '556be984-3fe7-4047-9b17-f6b0e940d03f', '2024-09-10', '2024-09-30', '2024-09-25'),
('0e51a616-98b9-41e1-a30d-54ac09d90fc7', 'db018be1-75b3-4ef3-b59c-dcd3fd44232c', '556be984-3fe7-4047-9b17-f6b0e940d03f', '2024-09-20', '2024-10-10', '2024-10-10'),
('0f4b760e-fd05-402a-b43c-225eb706f35b', 'c14ae712-f543-46b3-bd16-9fb3bf799b97', 'bb011bcc-1905-4c90-9d07-f7f8ee1a54fa', '2024-09-19', NULL, '2024-10-09'),
('2d883c0e-07d6-40cd-b4be-d91a1c37d5bc', 'db018be1-75b3-4ef3-b59c-dcd3fd44232c', '556be984-3fe7-4047-9b17-f6b0e940d03f', '2024-09-07', '2024-09-27', '2024-09-21'),
('5c0c3b42-042f-45f3-90b4-50d020194056', 'db018be1-75b3-4ef3-b59c-dcd3fd44232c', '556be984-3fe7-4047-9b17-f6b0e940d03f', '2024-09-13', '2024-10-03', '2024-10-03'),
('64358cca-30fa-4647-9325-5876a123509b', 'db018be1-75b3-4ef3-b59c-dcd3fd44232c', '556be984-3fe7-4047-9b17-f6b0e940d03f', '2024-09-12', '2024-10-02', '2024-10-02'),
('756a88f6-9669-435e-9313-b844fa4de9ba', 'db018be1-75b3-4ef3-b59c-dcd3fd44232c', '556be984-3fe7-4047-9b17-f6b0e940d03f', '2024-10-14', '2024-10-14', '2024-10-28'),
('fc2d588b-1547-4dc2-981f-2b3e4c6e14e0', 'db018be1-75b3-4ef3-b59c-dcd3fd44232c', '556be984-3fe7-4047-9b17-f6b0e940d03f', '2024-10-14', '2024-10-14', '2024-10-28');

INSERT INTO "public"."copies" ("id", "book_id", "acquired_date", "edition", "condition") VALUES
('066ed1e6-6d87-4f36-a8e0-730cca1092f8', 'e5d5713d-35fb-4981-b7ae-ed1db45723a5', '2021-09-10', '1st Edition', 'New'),
('0d3f84b2-88ba-458e-9156-f37cf1b00394', '26f85cce-b193-4da1-bc96-2e2c588604de', '2020-05-12', '3rd Edition', 'Good'),
('1aa719b7-dbcd-4722-988b-ec964685aa23', '4683079b-fb7d-410f-9ef2-1412185d4108', '2019-04-02', '3rd Edition', 'New'),
('2300fcfd-a4a0-49fc-bbbf-c76e4bf474b5', '26f85cce-b193-4da1-bc96-2e2c588604de', '2021-07-15', '3rd Edition', 'New'),
('2874f42a-623d-4a00-8977-9c05f4dea53c', 'ba0a97ff-6cae-4ba9-b6a4-cb5b3e386c2d', '2020-04-11', '3rd Edition', 'New'),
('36d86f15-edac-455f-bfca-9be8f804670d', '1aac4b04-035e-40b6-890c-6bdb735425ab', '2022-07-20', '2nd Edition', 'Fair'),
('424d4832-58c4-4857-a471-458b568a3ea8', '7dd2ca88-4ecb-4d45-a874-7591ce455708', '2019-10-01', '1st Edition', 'Fair'),
('5045ea5c-cd35-4834-9dc5-38e668bca5a9', '9f0161e3-42d4-4122-892b-dd466ff54225', '2021-03-12', '2nd Edition', 'Good'),
('890da60e-8a60-4fee-89a6-7c86932a29e2', 'e9b0de0d-8fa3-4bb2-abd0-6be71f9a4040', '2023-02-28', '1st Edition', 'New'),
('8dd153f4-65c7-4a5a-9bd4-72f28214438d', '10953403-56f8-4c6b-9926-b7f43a2bb35a', '2020-06-19', '2nd Edition', 'Good'),
('91a6b651-0fe7-42cc-bd44-7ea34ce10ce4', '22aee4e5-a8f5-43da-8af4-6386ac990798', '2022-03-10', '1st Edition', 'Fair'),
('95578b3f-8ff1-4126-b97e-eb2394b55751', '9f0161e3-42d4-4122-892b-dd466ff54225', '2020-08-14', '1st Edition', 'New'),
('a6694678-c496-4bf9-af79-b5f8d988b766', 'e1e0f7d9-d17f-454d-9446-6834c804ba46', '2022-05-18', '2nd Edition', 'New'),
('b2a7b5d4-97f1-4e72-b0ab-fcffcee591ee', '10953403-56f8-4c6b-9926-b7f43a2bb35a', '2019-11-22', '1st Edition', 'Fair'),
('b67f5ed2-3cab-46bb-a933-84d1b39383aa', '91c54736-2882-422e-b322-5365722c1641', '2021-12-15', '1st Edition', 'Good'),
('c14ae712-f543-46b3-bd16-9fb3bf799b97', '52096833-d586-4846-ab4c-450eb9dda592', '2023-01-15', '1st Edition', 'New'),
('c57614b4-3727-4caf-bcbf-9794debf416c', '4683079b-fb7d-410f-9ef2-1412185d4108', '2018-06-10', '2nd Edition', 'Good'),
('db018be1-75b3-4ef3-b59c-dcd3fd44232c', '52096833-d586-4846-ab4c-450eb9dda592', '2023-01-16', '1st Edition', 'Good'),
('e47092a2-6737-40f3-9774-cea567760d10', 'e1e0f7d9-d17f-454d-9446-6834c804ba46', '2021-09-23', '2nd Edition', 'Good');

INSERT INTO "public"."users" ("id", "full_name", "email", "phone_number", "registered_date") VALUES
('556be984-3fe7-4047-9b17-f6b0e940d03f', 'Jane Smith', 'jane.smith@example.com', '555-5678', '2024-10-13'),
('7da5281e-f6b1-41e4-856c-ebd04b3d5a9f', 'Don Knotts', 'don.knots@example.com', '555-1234', '2024-10-13'),
('bb011bcc-1905-4c90-9d07-f7f8ee1a54fa', 'John Doe', 'john.doe@example.com', '555-1234', '2024-10-13');