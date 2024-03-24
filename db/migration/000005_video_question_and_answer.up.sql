CREATE TABLE IF NOT EXISTS video_question
(
    id serial PRIMARY KEY,
    question character varying(300)
);

CREATE TABLE IF NOT EXISTS question_answer
(
    video_question_id integer,
    user_id integer,
    video_path character varying(400),
    PRIMARY KEY (video_question_id, user_id),
	FOREIGN KEY (video_question_id) REFERENCES public.video_question (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION,
	FOREIGN KEY (user_id)
    REFERENCES public.users (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
);