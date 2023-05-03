--
-- PostgreSQL database dump
--

-- Dumped from database version 14.1
-- Dumped by pg_dump version 15.2

-- Started on 2023-05-03 12:00:39 +05

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
-- TOC entry 3313 (class 1262 OID 16384)
-- Name: sdn; Type: DATABASE; Schema: -; Owner: -
--

CREATE DATABASE sdn WITH ENCODING = 'UTF8' LOCALE = 'en_US.utf8';


\connect sdn

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
-- Name: public; Type: SCHEMA; Schema: -; Owner: -
--

-- *not* creating schema, since initdb creates it


--
-- TOC entry 222 (class 1255 OID 16493)
-- Name: state_f(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.state_f() RETURNS integer
    LANGUAGE plpgsql
    AS $$BEGIN
  RETURN (SELECT COUNT(*)::int FROM public.list);
END;$$;


--
-- TOC entry 221 (class 1255 OID 16408)
-- Name: update_f(jsonb); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.update_f(jsonobject jsonb) RETURNS integer
    LANGUAGE plpgsql
    AS $$BEGIN
    
	DELETE FROM public.list;
 
    INSERT INTO public.list (
	uid, title, first_name, last_name, remarks, 
	program_list, aka_list, address_list, id_list, date_of_birth_list, 
	place_of_birth_list, nationality_list, citizenship_list, vessel_info)
	SELECT  
		CAST(i::json ->> 'uid' as integer), i::json ->> 'title', i::json ->> 'firstName', i::json ->> 'lastName', 
		i::json ->> 'remarks', CAST(i::json ->> 'programList' as json), CAST(i::json ->> 'akaList' as json), 
		CAST(i::json ->> 'addressList' as json), CAST(i::json ->> 'idList' as json), CAST(i::json ->> 'dateOfBirthList' as json), CAST(i::json ->> 'placeOfBirthList' as json), 
		CAST(i::json ->> 'nationalityList' as json), CAST(i::json ->> 'citizenshipList' as json), CAST(i::json ->> 'vesselInfo' as jsonb)
	FROM jsonb_array_elements(jsonobject -> 'sdnEntry') i;

RETURN 0;

EXCEPTION WHEN OTHERS THEN
  RAISE EXCEPTION '% %', SQLERRM, SQLSTATE;
END;
$$;


SET default_table_access_method = heap;

--
-- TOC entry 209 (class 1259 OID 16385)
-- Name: list; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.list (
    uid bigint NOT NULL,
    title character varying(1024),
    first_name character(100),
    last_name character varying(100) NOT NULL,
    remarks character varying(1024),
    program_list json,
    aka_list json,
    address_list json,
    id_list json,
    date_of_birth_list json,
    place_of_birth_list json,
    nationality_list json,
    citizenship_list json,
    vessel_info jsonb
);


--
-- TOC entry 3168 (class 2606 OID 16391)
-- Name: list list_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.list
    ADD CONSTRAINT list_pkey PRIMARY KEY (uid);


-- Completed on 2023-05-03 12:00:40 +05

--
-- PostgreSQL database dump complete
--

