INSERT INTO
    usuarios (nome, nick, email, senha)
VALUES
    (
        "Usuário 1",
        "usuario_1",
        "usuario1@gmail.com",
        "$2a$10$0iGYlKCAYTyJV/vC6nLGgeWFwD6AhSkWLsVRO/.M4lNK8OtIkfggy"
    ),
    (
        "Usuário 2",
        "usuario_2",
        "usuario2@gmail.com",
        "$2a$10$0iGYlKCAYTyJV/vC6nLGgeWFwD6AhSkWLsVRO/.M4lNK8OtIkfggy"
    ),
    (
        "Usuário 3",
        "usuario_3",
        "usuario3@gmail.com",
        "$2a$10$0iGYlKCAYTyJV/vC6nLGgeWFwD6AhSkWLsVRO/.M4lNK8OtIkfggy"
    );

INSERT INTO
    seguidores(usuario_id, seguidor_id)
VALUES
    (1, 2),
    (3, 1),
    (1, 3);