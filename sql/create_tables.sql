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
    varietalId INT NOT NULL,
    FarmerId INT NOT NULL,
    OriginId INT NOT NULL,
    ProccessId INT NOT NULL,
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
    altitude INT
);
-- Create table 'Origin'
CREATE TABLE public.origin (
    originid INT PRIMARY KEY,
    country CHAR(255) NOT NULL,
    state CHAR(255),
    citytown CHAR(255)
);
-- Create table 'Proccess'
CREATE TABLE public.proccess (
    proccessid INT PRIMARY KEY,
    name CHAR(255) NOT NULL,
    description CHAR(255),
    timeduration CHAR(255)
);
-- Create table 'Prices'
CREATE TABLE public.prices (
    priceid INT PRIMARY KEY,
    coffeeid INT NOT NULL,
    grams INT NOT NULL,
    values DECIMAL(10, 2) NOT NULL
);



-- Insert data into table 'Farmers'
INSERT INTO public.farmers (farmerid, name, description, altitude) VALUES
(1, 'Farmer 1', 'This is farmer 1', 1000),
(2, 'Farmer 2', 'This is farmer 2', 1500),
(3, 'Farmer 3', 'This is farmer 3', 2000);



-- Insert data into table 'Origin'
INSERT INTO public.origin (originid, country, state, citytown) VALUES
(1, 'Colombia', 'Antioquia', 'Medellín'),
(2, 'Brazil', 'Minas Gerais', 'Belo Horizonte'),
(3, 'Ethiopia', 'Sidamo', 'Yirgacheffe'),
(4, 'Colombia', 'Huila', 'Tres Esquinas');



-- Insert data into table 'Proccess'
INSERT INTO public.proccess (proccessid, name, description, timeduration) VALUES
(1, 'Washed', 'This is the washed process', '24 hours'),
(2, 'Natural', 'This is the natural process', '10 days'),
(3, 'Honey', 'This is the honey process', '30 days');


-- Insert data into table 'Prices'
INSERT INTO public.prices (priceid, coffeeid, grams, values) VALUES
(1, 1, 250, 10.00),
(2, 2, 500, 15.00),
(3, 3, 1000, 20.00);



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
    varietalId,
    FarmerId,
    OriginId,
    ProccessId,
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
    ('Café de Altura', 1, 1, 2, 1, 86.5, 'High', 'Medium', 'Good', 'Clear', 'Sweet', 1, 'https://www.antesuncafe.com/static/media/project-img1.42b4985dc6e22b4b5394.png', '2024-06-01 00:00:00', '2024-06-01 00:00:00'),
    ('Café Excelso', 2, 2, 3, 2, 85.0, 'Medium', 'Full', 'Good', 'Clear', 'Sweet', 2, 'https://www.antesuncafe.com/static/media/project-img2.8ccdcab64328946462b4.png','2024-06-01 00:00:00', '2024-06-01 00:00:00'),
    ('Café Supremo', 3, 3, 3, 3, 83.5, 'Low', 'Medium', 'Fair', 'Clear', 'Sweet', 3, 'https://www.antesuncafe.com/static/media/project-img1.42b4985dc6e22b4b5394.png', '2024-06-01 00:00:00', '2024-06-01 00:00:00'),
    ('Café Gourmet', 2, 2, 4, 3, 82.0, 'High', 'Full', 'Good', 'Clear', 'Sweet', 1 , 'https://www.antesuncafe.com/static/media/project-img1.42b4985dc6e22b4b5394.png', '2024-06-01 00:00:00', '2024-06-01 00:00:00'),
    ('Café Colombia', 1, 1, 2, 1, 86.5, 'High', 'Medium', 'Good', 'Clear', 'Sweet', 1, 'https://www.antesuncafe.com/static/media/project-img2.8ccdcab64328946462b4.png', '2024-06-01 00:00:00', '2024-06-01 00:00:00')
    ('Café Tabi', 1, 1, 2, 1, 88.5, 'High', 'Medium', 'Good', 'Clear', 'Sweet', 1, 'https://www.antesuncafe.com/static/media/project-img2.8ccdcab64328946462b4.png', '2024-06-01 00:00:00', '2024-06-01 00:00:00');
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
