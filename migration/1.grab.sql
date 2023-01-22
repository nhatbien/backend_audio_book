

DROP TABLE public.user_role;
DROP TABLE public.role_permission;
DROP TABLE public.users;
DROP TABLE public.role;
DROP TABLE public.permission;

DELETE FROM public.user_role;
DELETE FROM public.users;



CREATE TABLE public.role
(
    id serial NOT NULL,
    role_name text NOT NULL UNIQUE,
    role_description text,
    PRIMARY KEY (id)
);



INSERT INTO public.role (role_Name, role_description)
VALUES ('MEMBER', 'Tài khoản thường'),('PARTNER', 'Tài khoản Partner'),('ADMIN', 'Tài khoản Admin');



CREATE TABLE public.users
(
    id text NOT NULL,
    username text NOT NULL UNIQUE,
    email text UNIQUE,
    phone text NOT NULL UNIQUE,
    password text NOT NULL,
    full_name text,
    age bigint,
    address text,
    photo text,
    status bigint,
    role_id bigint NOT NULL,
    created_at date,
    updated_at date,
    PRIMARY KEY (id),
    FOREIGN KEY (role_id) REFERENCES public.role(id)

);





CREATE TABLE public.permission
(
   id serial,
   permission_name text UNIQUE,
   PRIMARY KEY (id)
);

CREATE TABLE public.role_permission
(
   role_id bigint references public.role(id),
   permission_id bigint references public.permission(id),
   PRIMARY KEY (permission_id, role_id)
);



