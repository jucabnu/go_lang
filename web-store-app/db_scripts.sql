-- Rodar esse script no console do postgres

CREATE DATABASE juliano_loja;

CREATE TABLE produtos (
    id serial primary key,
    nome varchar,
    descricao varchar, 
    preco decimal
    quantidade integer
);

INSERT INTO produtos (nome, descricao, preco, quantidade) VALUES
    ('Camiseta', 'Azul', 29.9, 10),
    ('Tênis', 'Baixo', 69, 4),
    ('Fones', 'Muito bons', 199.54, 3);
