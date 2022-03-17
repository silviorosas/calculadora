
CREATE TABLE localidades
(
  id          INT       NOT NULL AUTO_INCREMENT,
  created_at  timestamp NULL    ,
  updated_at  timestamp NULL    ,
  descripcion TEXT      NULL    ,
  PRIMARY KEY (id)
);

CREATE TABLE matriculas
(
  id             INT       NOT NULL AUTO_INCREMENT,
  created_at     timestamp NULL    ,
  updated_at     timestamp NULL    ,
  nombre         TEXT      NULL    ,
  apellido       TEXT      NULL    ,
  correo         TEXT      NULL    ,
  dni            INT       NOT NULL,
  matricula      INT       NOT NULL,
  resolucion     INT       NULL    ,
  telefono       INT       NULL    ,
  direccion      TEXT      NULL    ,
  url            TEXT      NULL    ,
  facebook       TEXT      NULL    ,
  linkedln       TEXT      NULL    ,
  aptitudes      TEXT      NULL    ,
  observaciones  TEXT      NULL    ,
  id_localidades INT       NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE titulos
(
  id               INT       NOT NULL AUTO_INCREMENT,
  created_at       timestamp NULL    ,
  updated_at       timestamp NULL    ,
  descripcion      TEXT      NOT NULL,
  id_universidades INT       NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE titulos_matriculas
(
  id            INT       NOT NULL AUTO_INCREMENT,
  created_at    timestamp NULL    ,
  updated_at    timestamp NULL    ,
  año_promocion TIMESTAMP NOT NULL,
  principal     BOOLEAN   NULL     DEFAULT false,
  id_matriculas INT       NOT NULL,
  id_titulos    INT       NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE universidades
(
  id          INT       NOT NULL AUTO_INCREMENT,
  created_at  timestamp NULL    ,
  updated_at  timestamp NULL    ,
  descripcion TEXT      NOT NULL,
  PRIMARY KEY (id)
);

ALTER TABLE titulos
  ADD CONSTRAINT FK_universidades_TO_titulos
    FOREIGN KEY (id_universidades)
    REFERENCES universidades (id);

ALTER TABLE matriculas
  ADD CONSTRAINT FK_localidades_TO_matriculas
    FOREIGN KEY (id_localidades)
    REFERENCES localidades (id);

ALTER TABLE titulos_matriculas
  ADD CONSTRAINT FK_matriculas_TO_titulos_matriculas
    FOREIGN KEY (id_matriculas)
    REFERENCES matriculas (id);

ALTER TABLE titulos_matriculas
  ADD CONSTRAINT FK_titulos_TO_titulos_matriculas
    FOREIGN KEY (id_titulos)
    REFERENCES titulos (id);

CREATE INDEX id
  ON titulos_matriculas (id ASC);

CREATE INDEX id_matriculas
  ON titulos_matriculas (id_matriculas ASC);

CREATE INDEX id
  ON matriculas (id ASC);

CREATE INDEX dni
  ON matriculas (dni ASC);

CREATE INDEX matricula
  ON matriculas (matricula ASC);
