--
-- PostgreSQL database dump
--

-- Dumped from database version 14.6 (Ubuntu 14.6-1.pgdg22.04+1)
-- Dumped by pg_dump version 15.1 (Ubuntu 15.1-1.pgdg22.04+1)

-- Started on 2022-12-15 18:48:07 WIB

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 4 (class 2615 OID 2200)
-- Name: public; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA public;


ALTER SCHEMA public OWNER TO postgres;

--
-- TOC entry 3401 (class 0 OID 0)
-- Dependencies: 4
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON SCHEMA public IS 'standard public schema';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 214 (class 1259 OID 16695)
-- Name: source_of_funds; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.source_of_funds (
    id integer NOT NULL,
    source character varying NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.source_of_funds OWNER TO postgres;

--
-- TOC entry 213 (class 1259 OID 16694)
-- Name: source_of_funds_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.source_of_funds_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.source_of_funds_id_seq OWNER TO postgres;

--
-- TOC entry 3403 (class 0 OID 0)
-- Dependencies: 213
-- Name: source_of_funds_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.source_of_funds_id_seq OWNED BY public.source_of_funds.id;


--
-- TOC entry 216 (class 1259 OID 16789)
-- Name: transactions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.transactions (
    id integer NOT NULL,
    sender_id integer,
    source_of_fund_id integer,
    recipient_id integer NOT NULL,
    amount double precision NOT NULL,
    description character varying,
    date timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.transactions OWNER TO postgres;

--
-- TOC entry 215 (class 1259 OID 16788)
-- Name: transactions_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.transactions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.transactions_id_seq OWNER TO postgres;

--
-- TOC entry 3404 (class 0 OID 0)
-- Dependencies: 215
-- Name: transactions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.transactions_id_seq OWNED BY public.transactions.id;


--
-- TOC entry 210 (class 1259 OID 16668)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id integer NOT NULL,
    email character varying NOT NULL,
    password character varying NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 209 (class 1259 OID 16667)
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO postgres;

--
-- TOC entry 3405 (class 0 OID 0)
-- Dependencies: 209
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- TOC entry 211 (class 1259 OID 16680)
-- Name: wallets_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.wallets_seq
    START WITH 157001
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.wallets_seq OWNER TO postgres;

--
-- TOC entry 212 (class 1259 OID 16681)
-- Name: wallets; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.wallets (
    id integer DEFAULT nextval('public.wallets_seq'::regclass) NOT NULL,
    user_id integer NOT NULL,
    balance double precision NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.wallets OWNER TO postgres;

--
-- TOC entry 3228 (class 2604 OID 16698)
-- Name: source_of_funds id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.source_of_funds ALTER COLUMN id SET DEFAULT nextval('public.source_of_funds_id_seq'::regclass);


--
-- TOC entry 3231 (class 2604 OID 16792)
-- Name: transactions id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions ALTER COLUMN id SET DEFAULT nextval('public.transactions_id_seq'::regclass);


--
-- TOC entry 3222 (class 2604 OID 16671)
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- TOC entry 3393 (class 0 OID 16695)
-- Dependencies: 214
-- Data for Name: source_of_funds; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.source_of_funds (id, source, created_at, updated_at, deleted_at) FROM stdin;
1	Bank Transfer	2022-12-14 14:33:40.139552	2022-12-14 14:33:40.139552	\N
2	Credit Card	2022-12-14 14:33:40.139552	2022-12-14 14:33:40.139552	\N
3	Cash	2022-12-14 14:33:40.139552	2022-12-14 14:33:40.139552	\N
\.


--
-- TOC entry 3395 (class 0 OID 16789)
-- Dependencies: 216
-- Data for Name: transactions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.transactions (id, sender_id, source_of_fund_id, recipient_id, amount, description, date, created_at, updated_at, deleted_at) FROM stdin;
20	157004	\N	157003	6500	tes transfer	2022-12-15 18:37:49.230037	2022-12-15 18:37:49.237367	2022-12-15 18:37:49.237367	\N
1	\N	1	157001	1000000	Top Up from Bank Transfer	2020-12-14 15:41:25.633	2022-12-14 15:41:25.645157	2022-12-14 15:41:25.645157	\N
2	157001	\N	157002	1000000	tes	2021-10-04 16:10:58.39	2022-12-14 16:10:58.404166	2022-12-14 16:10:58.404166	\N
3	157001	\N	157002	1000000	tes	2022-07-20 16:13:01.364	2022-12-14 16:13:01.37836	2022-12-14 16:13:01.37836	\N
4	157001	\N	157002	1000000	tes	2020-12-21 16:14:23.829	2022-12-14 16:14:23.841802	2022-12-14 16:14:23.841802	\N
5	157001	\N	157002	1000000	tes	2021-12-14 16:15:54.984	2022-12-14 16:15:55.002477	2022-12-14 16:15:55.002477	\N
6	157001	\N	157002	1000000	tes	2021-01-15 16:16:08.748	2022-12-14 16:16:08.756473	2022-12-14 16:16:08.756473	\N
7	157001	\N	157002	1000000	tes	2021-05-01 16:23:08.511	2022-12-14 16:23:08.517794	2022-12-14 16:23:08.517794	\N
8	157001	\N	157002	1000000	tes	2021-02-17 17:09:02.624	2022-12-14 17:09:02.636263	2022-12-14 17:09:02.636263	\N
9	\N	2	157001	1000000	Top Up from Credit Card	2021-06-03 18:29:59.663	2022-12-15 18:29:59.668893	2022-12-15 18:29:59.668893	\N
10	157001	\N	157002	1000000	tes transfer	2020-12-31 18:31:36.269	2022-12-15 18:31:36.277366	2022-12-15 18:31:36.277366	\N
11	\N	3	157002	1000000	Top Up from Cash	2021-01-01 18:32:42.916	2022-12-15 18:32:42.923754	2022-12-15 18:32:42.923754	\N
12	157002	\N	157005	1000000	tes transfer	2021-12-15 18:34:17.156	2022-12-15 18:34:17.164421	2022-12-15 18:34:17.164421	\N
13	\N	1	157002	2000000	Top Up from Bank Transfer	2021-04-23 18:34:31.829	2022-12-15 18:34:31.836283	2022-12-15 18:34:31.836283	\N
14	157002	\N	157004	1000000	tes transfer	2020-11-10 18:34:38.423	2022-12-15 18:34:38.431326	2022-12-15 18:34:38.431326	\N
15	\N	2	157003	5000000	Top Up from Credit Card	2021-08-09 18:35:53.643	2022-12-15 18:35:53.650465	2022-12-15 18:35:53.650465	\N
16	157003	\N	157001	1000000	tes transfer	2021-11-10 18:36:01.955	2022-12-15 18:36:01.963156	2022-12-15 18:36:01.963156	\N
17	157003	\N	157005	1000000	tes transfer	2021-01-15 18:36:16.637	2022-12-15 18:36:16.645019	2022-12-15 18:36:16.645019	\N
18	\N	3	157004	10000000	Top Up from Cash	2020-12-20 18:37:28.846	2022-12-15 18:37:28.853621	2022-12-15 18:37:28.853621	\N
19	157004	\N	157002	1000	tes transfer	2021-01-24 18:37:40.825	2022-12-15 18:37:40.833187	2022-12-15 18:37:40.833187	\N
\.


--
-- TOC entry 3389 (class 0 OID 16668)
-- Dependencies: 210
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, email, password, created_at, updated_at, deleted_at) FROM stdin;
1	andra@gmail.com	$2a$12$5M12TqOvjfiLLVmyXdd9fOt6t3bE0t7KGkAw7RnqXTFRWSiaTaSj.	2022-12-14 11:22:02.861221	2022-12-14 11:22:02.861221	\N
2	andra2@gmail.com	$2a$04$GW7W3ae9LuHpXePKZOJLbuwk788ToXB2G5hjwK1ZPVvNdp14gcpqu	2022-12-14 16:10:00.618595	2022-12-14 16:10:00.618595	\N
3	andra3@gmail.com	$2a$04$LbY3MSHyPnQmzkYSf5C2je0X1PZK.nRsPfyA3BZBWUqL6pQBt5Fiy	2022-12-15 18:26:12.234926	2022-12-15 18:26:12.234926	\N
4	andra4@gmail.com	$2a$04$9i0bTkxTAbyIVlv.FYPgwOnahKV2ICGNvvgr5QO6MTHycaogb/RC2	2022-12-15 18:26:16.887347	2022-12-15 18:26:16.887347	\N
5	andra5@gmail.com	$2a$04$dNvte9hq1HsoWF43BhYDA.iOOAyNAauuGIfF/WO8u7GE1PaczZkAO	2022-12-15 18:26:21.309999	2022-12-15 18:26:21.309999	\N
\.


--
-- TOC entry 3391 (class 0 OID 16681)
-- Dependencies: 212
-- Data for Name: wallets; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.wallets (id, user_id, balance, created_at, updated_at, deleted_at) FROM stdin;
157001	1	1000000	2022-12-14 11:22:02.862004	2022-12-15 18:36:01.95559	\N
157005	5	2000000	2022-12-15 18:26:21.310233	2022-12-15 18:36:16.63733	\N
157002	2	10001000	2022-12-14 16:10:00.619355	2022-12-15 18:37:40.825691	\N
157003	3	4006500	2022-12-15 18:26:12.241674	2022-12-15 18:37:49.23025	\N
157004	4	10992500	2022-12-15 18:26:16.88761	2022-12-15 18:37:49.236977	\N
\.


--
-- TOC entry 3406 (class 0 OID 0)
-- Dependencies: 213
-- Name: source_of_funds_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.source_of_funds_id_seq', 3, true);


--
-- TOC entry 3407 (class 0 OID 0)
-- Dependencies: 215
-- Name: transactions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.transactions_id_seq', 20, true);


--
-- TOC entry 3408 (class 0 OID 0)
-- Dependencies: 209
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 5, true);


--
-- TOC entry 3409 (class 0 OID 0)
-- Dependencies: 211
-- Name: wallets_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.wallets_seq', 157005, true);


--
-- TOC entry 3242 (class 2606 OID 16704)
-- Name: source_of_funds sofs_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.source_of_funds
    ADD CONSTRAINT sofs_pkey PRIMARY KEY (id);


--
-- TOC entry 3244 (class 2606 OID 16799)
-- Name: transactions transactions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_pkey PRIMARY KEY (id);


--
-- TOC entry 3236 (class 2606 OID 16677)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- TOC entry 3238 (class 2606 OID 16679)
-- Name: users users_un; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_un UNIQUE (email);


--
-- TOC entry 3240 (class 2606 OID 16688)
-- Name: wallets wallets_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wallets
    ADD CONSTRAINT wallets_pkey PRIMARY KEY (id);


--
-- TOC entry 3247 (class 2606 OID 16800)
-- Name: transactions senders_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT senders_fkey FOREIGN KEY (sender_id) REFERENCES public.wallets(id);


--
-- TOC entry 3248 (class 2606 OID 16805)
-- Name: transactions sofs_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT sofs_fkey FOREIGN KEY (source_of_fund_id) REFERENCES public.source_of_funds(id);


--
-- TOC entry 3245 (class 2606 OID 16689)
-- Name: wallets users_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wallets
    ADD CONSTRAINT users_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- TOC entry 3246 (class 2606 OID 16810)
-- Name: wallets wallets_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wallets
    ADD CONSTRAINT wallets_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- TOC entry 3402 (class 0 OID 0)
-- Dependencies: 4
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE USAGE ON SCHEMA public FROM PUBLIC;
GRANT ALL ON SCHEMA public TO PUBLIC;


-- Completed on 2022-12-15 18:48:07 WIB

--
-- PostgreSQL database dump complete
--

