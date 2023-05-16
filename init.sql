CREATE TABLE IF NOT EXISTS users(
                                    id SERIAL NOT NULL UNIQUE,
                                    name VARCHAR(255) NOT NULL,
                                    email VARCHAR(255) NOT NULL,
                                    password_hash VARCHAR(255),
                                    registered_at DATE NOT NULL,
                                    user_type VARCHAR(20)
);

CREATE TABLE IF NOT EXISTS agents(
                                    id SERIAL NOT NULL UNIQUE,
                                    user_id INT REFERENCES users(id) ON DELETE CASCADE NOT NULL
);

CREATE TABLE IF NOT EXISTS owners(
                                    id SERIAL PRIMARY KEY UNIQUE NOT NULL,
                                    user_id INT REFERENCES users(id) ON DELETE CASCADE NOT NULL
);

CREATE TABLE IF NOT EXISTS clients(
                                     id SERIAL NOT NULL UNIQUE,
                                     user_id INT REFERENCES users(id) ON DELETE CASCADE NOT NULL
);

CREATE TABLE IF NOT EXISTS houses(
                                     id SERIAL NOT NULL UNIQUE,
                                     address VARCHAR(255) NOT NULL,
                                     owner_id INT REFERENCES owners(id) ON DELETE CASCADE NOT NULL,
                                     agent_id INT REFERENCES agents(id) ON DELETE CASCADE,
                                     build_date DATE,
                                     price NUMERIC(18, 2)
);

CREATE TABLE IF NOT EXISTS client_cart(
                                    client_id INT REFERENCES clients(id) ON DELETE CASCADE NOT NULL,
                                    house_id INT REFERENCES houses(id) ON DELETE CASCADE NOT NULL
);


