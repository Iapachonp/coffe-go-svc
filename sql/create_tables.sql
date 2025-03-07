--
-- PostgreSQL database dump
--

CREATE TABLE public.users (
    id INT NOT NULL,
    first_name CHAR(255),
    last_name CHAR(255),
    email CHAR(255),
    password CHAR(255),
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);


CREATE TABLE public.coffees (
    CoffeeId INT PRIMARY KEY,
    Name CHAR(255) NOT NULL,
    Description CHAR(255),
    varietalId INT NOT NULL,
    FarmerId INT NOT NULL,
    OriginId INT NOT NULL,
    processId INT NOT NULL,
    SCA DECIMAL(3, 1) NOT NULL,
    Acidity CHAR(255) NOT NULL,
    Body CHAR(255) NOT NULL,
    Balance CHAR(255) NOT NULL,
    Clarity CHAR(255) NOT NULL,
    Sweetness CHAR(255) NOT NULL,
    PriceId INT NOT NULL,
    Image CHAR(255), 
    CreatedAt timestamp without time zone,
    UpdatedAt timestamp without time zone
);

--
-- Name: genres_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

ALTER TABLE public.coffees ALTER COLUMN CoffeeId ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.coffees_CoffeeId_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: movies; Type: TABLE; Schema: public; Owner: -
--

-- Create the public.public.varietals table
CREATE TABLE public.varietals (
    varietalId INT PRIMARY KEY,
    Name CHAR(50) NOT NULL,
    Description CHAR(255),
    BeanSize CHAR(50),
    Stature CHAR(50),
    QualityPotential CHAR(50),
    OptimalAltitude CHAR(50)
);


ALTER TABLE public.varietals ALTER COLUMN varietalId ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.varietal_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

ALTER TABLE public.users ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);

-- Create table 'Farmers'
CREATE TABLE public.farmers (
    farmerid INT PRIMARY KEY,
    name CHAR(255) NOT NULL,
    description CHAR(255),
    altitude CHAR(255)
);

ALTER TABLE public.farmers ALTER COLUMN farmerId ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.farmer_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


-- Create table 'Origin'
CREATE TABLE public.origins (
    originid INT PRIMARY KEY,
    country CHAR(255) NOT NULL,
    state CHAR(255),
    citytown CHAR(255)
);


ALTER TABLE public.origins ALTER COLUMN originId ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.origin_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);

-- Create table 'process'
CREATE TABLE public.processes (
    processid INT PRIMARY KEY,
    name CHAR(255) NOT NULL,
    description CHAR(255),
    timeduration CHAR(255)
);

ALTER TABLE public.processes ALTER COLUMN processId ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.process_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);

-- Create table 'Prices'
CREATE TABLE public.prices (
    Priceid INT PRIMARY KEY,
    prices JSONB
);


ALTER TABLE public.prices ALTER COLUMN priceId ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.price_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


-- Insert data into table 'Prices'
INSERT INTO public.prices ( prices) VALUES
('[{ "grams": "250", "price" :"20000"} , {"grams": "500", "price" : "38000"} , {"grams": "1000", "price": "72000"} , {"grams":"2000", "price" :"170000" }]'),
('[{"grams": "500", "price" : "38000"} , {"grams": "1000", "price": "72000"} , {"grams":"2000", "price" :"170000" }]'),
('[{ "grams": "250", "price" :"38000"} , {"grams": "500", "price" : "60000"} , {"grams": "1000", "price": "110000"} , {"grams":"2000", "price" :"250000" }]'),
('[{ "grams": "250", "price" :"38000"} , {"grams": "500", "price" : "60000"}]'),
('[{ "grams": "250", "price" :"59000"} , {"grams": "1000", "price" : "99000"}]');

-- select prices -> '500gr' value from prices where priceid = 2;

-- select jsonb_path_query(prices, '$.grams'), jsonb_path_query(prices, '$.price')  from prices where priceid = 3;


INSERT INTO public.varietals (Name, Description, BeanSize, Stature, QualityPotential, OptimalAltitude) VALUES
('Bourbón', 'A Colombian coffee varietal known for its sweet, well-balanced flavour.', 'Medium', 'Tall', 'Excellent', '1200-1800m'),
('Caturra', 'A Colombian coffee varietal that is known for its high yield and resistance to disease.', 'Small', 'Short', 'Good', '1000-1600m'),
('Castillo', 'A Colombian coffee varietal that is a hybrid of Caturra and Typica.', 'Medium', 'Tall', 'Very Good', '1200-1800m'),
('Colombiana', 'A Colombian coffee varietal that is known for its large beans and rich flavour.', 'Large', 'Tall', 'Excellent', '1000-1600m'),
('Excelso', 'A Colombian coffee varietal that is known for its high quality and complex flavour.', 'Medium', 'Tall', 'Excellent', '1200-1800m'),
('Maragogype', 'A Colombian coffee varietal that is known for its giant beans and mild flavour.', 'Giant', 'Tall', 'Good', '1000-1600m'),
('Pache', 'A Colombian coffee varietal that is known for its unique, nutty flavour.', 'Small', 'Short', 'Good', '1200-1800m'),
('Tabi', 'A Colombian coffee varietal that is known for its high yield and resistance to disease.', 'Medium', 'Short', 'Good', '1000-1600m'),
('Typica', 'A Colombian coffee varietal that is known for its classic, balanced flavour.', 'Medium', 'Tall', 'Very Good', '1200-1800m'),
('Villa Sarchi', 'A Colombian coffee varietal that is known for its sweet, fruity flavour.', 'Small', 'Short', 'Good', '1000-1600m');



INSERT INTO public.coffees (
    Name,
    Description,
    varietalId,
    FarmerId,
    OriginId,
    processId,
    SCA,
    Acidity,
    Body,
    Balance,
    Clarity,
    Sweetness,
    PriceId,
    Image,
    Createdat,
    Updatedat
)
VALUES
    ('Café de Altura', 'Cafe de altura super util agregar algo aca',1, 1, 2, 1, 86.5, 'High', 'Medium', 'Good', 'Clear', 'Sweet', 1, 'https://www.antesuncafe.com/static/media/project-img1.42b4985dc6e22b4b5394.png', '2024-06-01 00:00:00', '2024-06-01 00:00:00'),
    ('Café Excelso', 'Cafe Excelso super util agregar algo aca', 2, 2, 3, 2, 85.0, 'Medium', 'Full', 'Good', 'Clear', 'Sweet', 2, 'https://www.antesuncafe.com/static/media/project-img2.8ccdcab64328946462b4.png','2024-06-01 00:00:00', '2024-06-01 00:00:00'),
    ('Café Supremo', 'Cafe Supremo super util agregar algo aca', 3, 3, 3, 3, 83.5, 'Low', 'Medium', 'Fair', 'Clear', 'Sweet', 3, 'https://www.antesuncafe.com/static/media/project-img1.42b4985dc6e22b4b5394.png', '2024-06-01 00:00:00', '2024-06-01 00:00:00'),
    ('Café Gourmet', 'Cafe Gourmet super util agregar algo aca', 2, 2, 4, 3, 82.0, 'High', 'Full', 'Good', 'Clear', 'Sweet', 1 , 'https://www.antesuncafe.com/static/media/project-img1.42b4985dc6e22b4b5394.png', '2024-06-01 00:00:00', '2024-06-01 00:00:00'),
    ('Café Colombia', 'Cafe Colombia super util agregar algo aca', 1, 1, 2, 1, 86.5, 'High', 'Medium', 'Good', 'Clear', 'Sweet', 1, 'https://www.antesuncafe.com/static/media/project-img2.8ccdcab64328946462b4.png', '2024-06-01 00:00:00', '2024-06-01 00:00:00'),
    ('Café Tabi', 'Cafe Tabi super util agregar algo aca', 1, 1, 2, 1, 88.5, 'High', 'Medium', 'Good', 'Clear', 'Sweet', 1, 'https://www.antesuncafe.com/static/media/project-img2.8ccdcab64328946462b4.png', '2024-06-01 00:00:00', '2024-06-01 00:00:00');
    -- ('Café Especial', 5, 5, 6, 5, 80.5, 'Medium', 'Medium', 'Fair', 'Clear', 'Sweet', 5, ''),
    -- ('Café Premium', 6, 6, 7, 6, 79.0, 'Low', 'Full', 'Good', 'Clear', 'Sweet', 6, ''),
    -- ('Café Arábigo', 7, 7, 8, 7, 77.5, 'High', 'Medium', 'Fair', 'Clear', 'Sweet', 7, ''),
    -- ('Café Robusta', 8, 8, 9, 8, 76.0, 'Medium', 'Full', 'Good', 'Clear', 'Sweet', 8, ''),
    -- ('Café Caturra', 9, 9, 10, 9, 74.5, 'Low', 'Medium', 'Fair', 'Clear', 'Sweet', 9, ''),
    -- ('Café Catimor', 10, 10, 11, 10, 73.0, 'High', 'Full', 'Good', 'Clear', 'Sweet', 10, ''),
    -- ('Café Geisha', 11, 11, 12, 11, 71.5, 'Medium', 'Medium', 'Fair', 'Clear', 'Sweet', 11, ''),
    -- ('Café Bourbon', 12, 12, 13, 12, 70.0, 'Low', 'Full', 'Good', 'Clear', 'Sweet', 12, ''),
    -- ('Café Typica', 13, 13, 14, 13, 68.5, 'High', 'Medium', 'Fair', 'Clear', 'Sweet', 13, ''),
    -- ('Café Sidra', 14, 14, 15, 14, 67.0, 'Medium', 'Full', 'Good', 'Clear', 'Sweet', 14, ''),
    -- ('Café Mokka', 15, 15, 16, 15, 65.5, 'Low', 'Medium', 'Fair', 'Clear', 'Sweet', 15, ''),
    -- ('Café Harrar', 16, 16, 17, 16, 64.0, 'High', 'Full', 'Good', 'Clear', 'Sweet', 16, ''),
    -- ('Café Yirgacheffe', 17, 17, 18, 17, 62.5, 'Medium', 'Medium', 'Fair', 'Clear', 'Sweet', 17, ''),
    -- ('Café Jamaica Blue Mountain', 18, 18, 19, 18, 61.0, 'Low', 'Full', 'Good', 'Clear', 'Sweet', 18, ''),
    -- ('Café Kona', 19, 19, 20, 19, 59.5, 'High', 'Medium', 'Fair', 'Clear', 'Sweet', 19, ''),
    -- ('Café Kopi Luwak', 20, 20, 21, 20, 58.0, 'Medium', 'Full', 'Good', 'Clear', 'Sweet', 20, '');

-- Insert data into table 'Farmers'
INSERT INTO public.farmers (name, description, altitude) VALUES
('Farmer 1', 'This is farmer 1', '1000'),
('Farmer 2', 'This is farmer 2', '1500'),
('Farmer 3', 'This is farmer 3', '2000');

-- Insert data into table 'Origin'
INSERT INTO public.origins ( country, state, citytown) VALUES
('Colombia', 'Antioquia', 'Medellín'),
('Brazil', 'Minas Gerais', 'Belo Horizonte'),
('Ethiopia', 'Sidamo', 'Yirgacheffe'),
('Colombia', 'Huila', 'Tres Esquinas');

-- Insert data into table 'process'
INSERT INTO public.processes ( name, description, timeduration) VALUES
('Washed', 'This is the washed process', '24 hours'),
('Natural', 'This is the natural process', '10 days'),
('Honey', 'This is the honey process', '30 days');
