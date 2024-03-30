CREATE TABLE IF NOT EXISTS test_question
(
    id serial PRIMARY KEY,
    question character varying(300)
);

CREATE TABLE IF NOT EXISTS test_answer
(
    id serial PRIMARY KEY,
    test_question_id integer,
    answer character varying(300),
    FOREIGN KEY (test_question_id) REFERENCES public.test_question (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
);

CREATE TABLE IF NOT EXISTS test_user_answer
(
    test_question_id integer,
    user_id integer,
    test_answer_id integer,
    PRIMARY KEY (test_question_id, user_id),
	FOREIGN KEY (test_question_id) REFERENCES public.test_question (id) MATCH SIMPLE
    ON UPDATE NO ACTION ON DELETE NO ACTION,
	FOREIGN KEY (user_id)
    REFERENCES public.users (id) MATCH SIMPLE
    ON UPDATE NO ACTION ON DELETE NO ACTION,
    FOREIGN KEY (test_answer_id) REFERENCES public.test_answer (id) MATCH SIMPLE
    ON UPDATE NO ACTION ON DELETE NO ACTION
);