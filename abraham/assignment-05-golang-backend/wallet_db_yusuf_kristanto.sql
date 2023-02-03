CREATE DATABASE wallet_db_yusuf_kristanto;

\c wallet_db_yusuf_kristanto

CREATE TABLE "users" ("id" bigserial,"name" text,"email" text,"password" text,"created_at" timestamptz,"deleted_at" timestamptz,PRIMARY KEY ("id"));

CREATE TABLE "transactions" ("id" bigserial,"sender" bigint,"recipient" bigint,"amount" bigint,"description" text,"created_at" timestamptz,PRIMARY KEY ("id"));

CREATE TABLE "wallets" ("id" bigserial,"number" bigint,"balance" bigint,"user_id" bigint,PRIMARY KEY ("id"),CONSTRAINT "fk_wallets_user" FOREIGN KEY ("user_id") REFERENCES "users"("id") ON DELETE SET NULL ON UPDATE CASCADE);

CREATE TABLE "fund_sources" ("id" bigserial,"method" text,PRIMARY KEY ("id"));

INSERT INTO public.fund_sources(
	method)
	VALUES ('Bank Transfer');
INSERT INTO public.fund_sources(
	method)
	VALUES ('Credit Card');
INSERT INTO public.fund_sources(
	method)
	VALUES ('Cash');
	
INSERT INTO public.users(
	name, email, password, created_at)
	VALUES ('Lewis Hamilton', 'lewis.hamilton@mercedes.com', '$2a$14$20Bg96CkbyVnwPvVyX2vReXUIXxAre7u.ZbHlcIWDriBgzjTrpX6a', now());
INSERT INTO public.users(
	name, email, password, created_at)
	VALUES ('Sebastian Vettel', 'sebastian.vettel@redbull.com', '$2a$14$kCYTmrnpJzwsINPjV/9U2uvfujkcHTDnMY4zwWHVRFvUKojiZiVWa', now());
INSERT INTO public.users(
	name, email, password, created_at)
	VALUES ('Michael Schumacher', 'michael.schumacher@ferrari.com', '$2a$14$65XO7riXR/Ez7EW7Ogd89ec/O5KhN5SRF60TbM6Dl5Rj0PHF7/6ve', now());
INSERT INTO public.users(
	name, email, password, created_at)
	VALUES ('James Hunt', 'james.hunt@mclaren.com', '$2a$14$cn5LLL80OuElWsE3/OxpKu4Ue4uLkk6pDR2qPw6eaCz0SuCYPZT4S', now());
INSERT INTO public.users(
	name, email, password, created_at)
	VALUES ('Fernando Alonso', 'fernando.alonso@renault.com', '$2a$14$sOHAgtRxH8sF5MkdmCim6.MQyBLAUXIOCpuB7.kSbKjRfC0sb1qI.', now());
	
INSERT INTO public.wallets(
	"number", balance, user_id)
	VALUES (777001, 0, 1);
INSERT INTO public.wallets(
	"number", balance, user_id)
	VALUES (777002, 0, 2);
INSERT INTO public.wallets(
	"number", balance, user_id)
	VALUES (777003, 0, 3);
INSERT INTO public.wallets(
	"number", balance, user_id)
	VALUES (777004, 0, 4);
INSERT INTO public.wallets(
	"number", balance, user_id)
	VALUES (777005, 0, 5);

INSERT INTO public.transactions(
	sender, recipient, amount, description, created_at)
	VALUES (777001, 777001, 100000, 'Top up from Cash', '2020-07-01');
INSERT INTO public.transactions(
	sender, recipient, amount, description, created_at)
	VALUES (777002, 777002, 500000, 'Top up from Credit Card', '2020-08-01');
INSERT INTO public.transactions(
	sender, recipient, amount, description, created_at)
	VALUES (777003, 777003, 600000, 'Top up from Bank Transfer', '2020-08-01');
INSERT INTO public.transactions(
	sender, recipient, amount, description, created_at)
	VALUES (777004, 777004, 300000, 'Top up from Cash', '2020-09-01');
INSERT INTO public.transactions(
	sender, recipient, amount, description, created_at)
	VALUES (777005, 777005, 800000, 'Top up from Credit Card', '2020-10-01');
INSERT INTO public.transactions(
	sender, recipient, amount, description, created_at)
	VALUES (777001, 777005, 10000, 'Lunch', '2020-11-01');
INSERT INTO public.transactions(
	sender, recipient, amount, description, created_at)
	VALUES (777002, 777005, 10000, 'Lunch', '2020-11-02');
INSERT INTO public.transactions(
	sender, recipient, amount, description, created_at)
	VALUES (777003, 777005, 10000, 'Lunch', '2020-11-03');
INSERT INTO public.transactions(
	sender, recipient, amount, description, created_at)
	VALUES (777004, 777005, 10000, 'Lunch', '2020-11-04');
INSERT INTO public.transactions(
	sender, recipient, amount, description, created_at)
	VALUES (777001, 777002, 10000, 'Toy car', '2020-11-10');
INSERT INTO public.transactions(
	sender, recipient, amount, description, created_at)
	VALUES (777003, 777002, 10000, 'Toy dinosaur', '2020-11-11');
INSERT INTO public.transactions(
	sender, recipient, amount, description, created_at)
	VALUES (777004, 777002, 10000, 'Toy doll', '2020-11-12');
INSERT INTO public.transactions(
	sender, recipient, amount, description, created_at)
	VALUES (777005, 777003, 10000, 'Gift', '2021-01-01');
INSERT INTO public.transactions(
	sender, recipient, amount, description, created_at)
	VALUES (777005, 777001, 10000, 'Gift', '2021-02-03');
INSERT INTO public.transactions(
	sender, recipient, amount, description, created_at)
	VALUES (777005, 777002, 10000, 'Business', '2021-03-15');
INSERT INTO public.transactions(
	sender, recipient, amount, description, created_at)
	VALUES (777001, 777002, 10000, 'Business', '2021-03-18');
INSERT INTO public.transactions(
	sender, recipient, amount, description, created_at)
	VALUES (777004, 777002, 10000, 'Business', '2021-04-01');
INSERT INTO public.transactions(
	sender, recipient, amount, description, created_at)
	VALUES (777002, 777002, 100000, 'Top up from Cash', '2022-01-01');
INSERT INTO public.transactions(
	sender, recipient, amount, description, created_at)
	VALUES (777003, 777003, 100000, 'Top up from Cash', '2022-04-01');
INSERT INTO public.transactions(
	sender, recipient, amount, description, created_at)
	VALUES (777004, 777004, 100000, 'Top up from Cash', '2022-05-01');