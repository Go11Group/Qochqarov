CREATE TABLE "branch" (
    "id" UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    "title" VARCHAR(30) NOT NULL
);

CREATE TABLE "teacher" (
    "id" UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    "branch_id" UUID REFERENCES "branch" ("id"),
    "first_name" VARCHAR(30) NOT NULL,
    "last_name" VARCHAR(30) NOT NULL,
    "salary" NUMERIC
);

CREATE TABLE "asisstant_teacher" (
    "id" UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    "branch_id" UUID REFERENCES "branch" ("id"),
    "first_name" VARCHAR(30) NOT NULL,
    "last_name" VARCHAR(30) NOT NULL,
    "salary" NUMERIC
);


CREATE TABLE "course" (
    "id" UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    "name" VARCHAR(30) NOT NULL,
    "teacher_id" UUID REFERENCES "teacher" ("id"),
    "asisstant_teacher_id" UUID REFERENCES "asisstant_teacher" ("id"),
    "price" NUMERIC
);

CREATE TABLE "students" (
    "id" UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    "first_name" VARCHAR(30) NOT NULL,
    "last_name" VARCHAR(30) NOT NULL,
    "course_id" UUID REFERENCES "course" ("id"),
    "payed_sum" NUMERIC
);
