CREATE TABLE filmes (
    id SERIAL PRIMARY KEY,
    nome VARCHAR,
    genero VARCHAR,
    descricao TEXT,
    imagem_capa BYTEA
);

INSERT INTO filmes (nome, genero, descricao) VALUES
('Sexta Feira 13', 'Terror', 'Todo mundo morre no final'),
('Carros', 'Animação', 'Vrum Vrum katichauuuu');
