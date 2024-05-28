create table
    student (
        id uuid primary key default gen_random_uuid () not null,
        name varchar,
        age int
    )
create table
    curs (
        id uuid primary key default gen_random_uuid () not null,
        name varchar
    )
create table
    student_curs (
        id uuid primary key default gen_random_uuid () not null,
        student_id uuid referances student (id),
        curs_id uuid referances curs (id)
    )
create table
    gred (
        id uuid primary key default gen_random_uuid () not null,
        student_curs_id uuid referances student_curs (id),
        ball int
    )
INSERT INTO
    student (name, age)
VALUES
    ('John', 25),
    ('Alice', 22),
    ('Bob', 21),
    ('Emma', 23),
    ('Michael', 20),
    ('Sophia', 24),
    ('William', 22),
    ('Olivia', 21),
    ('James', 23),
    ('Charlotte', 20),
    ('Alexander', 22),
    ('Emily', 21),
    ('Daniel', 24),
    ('Ava', 23),
    ('Ethan', 20),
    ('Mia', 22),
    ('Matthew', 21),
    ('Abigail', 24),
    ('Benjamin', 23),
    ('Grace', 20);

INSERT INTO
    curs (name)
VALUES
    ('Mathematics'),
    ('Physics'),
    ('Chemistry'),
    ('Biology'),
    ('Computer Science'),
    ('History'),
    ('Literature'),
    ('Art'),
    ('Music'),
    ('Physical Education');

INSERT INTO
    student_curs (student_id, curs_id)
VALUES
    (
        'c8e65900-3318-4dd0-9f97-b1c78c4d99d5',
        'f360208e-7dce-4a80-b388-346a374e8ade'
    ), -- John - Mathematics
    (
        'c8e65900-3318-4dd0-9f97-b1c78c4d99d5',
        '5fe9fad2-78cd-4021-9540-4bf9008ebff7'
    ), -- John - Physics
    (
        '31925996-fefc-4579-9f26-9d6c666c2b5b',
        '09716a6b-786a-46f7-bb12-ddfa4265aa47'
    ), -- Alice - Chemistry
    (
        '31925996-fefc-4579-9f26-9d6c666c2b5b',
        'fbd92a8c-315f-4e74-93df-a52e1d905aef'
    ), -- Alice - Biology
    (
        '74bfd87e-993a-4f1e-91f2-fa62d7b22b55',
        '55aee8dd-622b-4d09-8dd0-dab52e6a4769'
    ), -- Bob - Computer Science
    (
        '74bfd87e-993a-4f1e-91f2-fa62d7b22b55',
        '716d2ec4-07f9-4dac-8e07-55d777b2a7bb'
    ), -- Bob - History
    (
        '487e8638-a089-4e79-b3df-6dd8759f958e',
        '9768362b-8b44-4f35-bbc6-307903f6e8ca'
    ), -- Emma - Literature
    (
        '487e8638-a089-4e79-b3df-6dd8759f958e',
        '08551c25-12d2-4d71-993d-d82ac9113f70'
    ), -- Emma - Art
    (
        '6ae0c5b9-eb90-42cf-a4a5-3074a4793471',
        '374a217f-25f1-4931-bf96-8e8ab531f344'
    ), -- Michael - Music
    (
        '6ae0c5b9-eb90-42cf-a4a5-3074a4793471',
        '82bd36de-dda3-40cc-bbcb-7e342781394d'
    ), -- Michael - Physical Education
    (
        '65b38309-7efd-404e-a0bc-3c67d585d2be',
        'f360208e-7dce-4a80-b388-346a374e8ade'
    ), -- Sophia - Mathematics
    (
        '65b38309-7efd-404e-a0bc-3c67d585d2be',
        '5fe9fad2-78cd-4021-9540-4bf9008ebff7'
    ), -- Sophia - Physics
    (
        '83e68025-ec90-488b-8403-9c11ff2db374',
        '09716a6b-786a-46f7-bb12-ddfa4265aa47'
    ), -- William - Chemistry
    (
        '83e68025-ec90-488b-8403-9c11ff2db374',
        'fbd92a8c-315f-4e74-93df-a52e1d905aef'
    ), -- William - Biology
    (
        '7e213310-e1e7-4a03-94b7-7685f33ccb34',
        '55aee8dd-622b-4d09-8dd0-dab52e6a4769'
    ), -- Olivia - Computer Science
    (
        '7e213310-e1e7-4a03-94b7-7685f33ccb34',
        '716d2ec4-07f9-4dac-8e07-55d777b2a7bb'
    ), -- Olivia - History
    (
        'c3ab7d14-67d7-4358-8e74-faf7946d97d5',
        '9768362b-8b44-4f35-bbc6-307903f6e8ca'
    ), -- James - Literature
    (
        'c3ab7d14-67d7-4358-8e74-faf7946d97d5',
        '08551c25-12d2-4d71-993d-d82ac9113f70'
    ), -- James - Art
    (
        'ed261c35-1be8-48ab-9421-07035133f4c8',
        '374a217f-25f1-4931-bf96-8e8ab531f344'
    ), -- Alexander - Music
    (
        'ed261c35-1be8-48ab-9421-07035133f4c8',
        '82bd36de-dda3-40cc-bbcb-7e342781394d'
    ), -- Alexander - Physical Education
    (
        'c6b4954f-ace6-480a-9ec6-3f4363473cde',
        'f360208e-7dce-4a80-b388-346a374e8ade'
    ), -- Emily - Mathematics
    (
        'c6b4954f-ace6-480a-9ec6-3f4363473cde',
        '5fe9fad2-78cd-4021-9540-4bf9008ebff7'
    ), -- Emily - Physics
    (
        '728508e5-28fd-4710-9bde-5e961f0534dd',
        '09716a6b-786a-46f7-bb12-ddfa4265aa47'
    ), -- Daniel - Chemistry
    (
        '728508e5-28fd-4710-9bde-5e961f0534dd',
        'fbd92a8c-315f-4e74-93df-a52e1d905aef'
    ), -- Daniel - Biology
    (
        '4b68c958-bbc9-43e0-b7e0-16bc9e3d51b0',
        '55aee8dd-622b-4d09-8dd0-dab52e6a4769'
    );

-- Ava - Computer Science
INSERT INTO
    gred (student_curs_id, ball)
VALUES
    ('0289d864-d003-4faf-92cb-8919ef528f3d', 85),
    ('779bfec3-8612-473b-9b6c-135f377d79dc', 90),
    ('a4892453-27bc-444f-9914-74370ad9ada6', 75),
    ('7493634b-e9b8-4130-976d-b6a72c8ee52d', 88),
    ('c3b6a99c-73b7-489c-bfd1-d6fb1984c0d3', 92),
    ('9bb212a7-af15-406b-b946-64631c6d313c', 78),
    ('d8f44633-882d-4389-8372-78e6a32d995f', 87),
    ('1f39e537-5732-412c-be90-57b178cae7b1', 93),
    ('c48d76fd-5832-4ac8-8b56-f0bb0c432dcf', 80),
    ('b06349cf-9a41-4d89-b9b7-d74409538511', 85),
    ('154c2ae6-b84b-4094-b997-c1623dd88755', 89),
    ('3ef04927-3a29-4431-b28a-65b36a4dfaf4', 91),
    ('00f37b51-3791-4a47-bcd8-658841a07e98', 76),
    ('abc55945-0da3-4379-a5d1-ec4c0b4155bf', 83),
    ('6c41eeff-574f-4006-bcf2-89a4d9ca6a6d', 85),
    ('077454bc-0fcd-4816-a39b-18bbbb7cb452', 90),
    ('bb4f4e39-ac57-4f9a-8793-dac35734243b', 75),
    ('7dd31efe-bce7-4fdb-a5f5-9be1a91c2966', 88),
    ('bb5afee2-08dd-4f01-b174-bafbb3a94b85', 92),
    ('628e99d3-f178-48cb-bd34-df0f5bfe9c1a', 78),
    ('6013cce8-df29-46ac-beaa-e995ba6db995', 87),
    ('36e6f8b9-6eeb-445e-b8a5-8035a7f35b11', 93),
    ('84ce5f14-cac2-44bd-a8d4-30a2c0b28877', 80),
    ('2e3d16c2-96fb-4569-a410-c6861c0e3169', 85),
    ('057df5b9-2f15-4cf2-af08-02ad2a83a875', 90);

--2-savol
select
    c.name,
    array_agg (s2.name),
    array_agg (g.ball)
from
    student_curs as s1
    join student as s2 on s1.student_id = s2.id
    join curs as c on curs_id = c.id
    join gred as g on g.student_curs_id = s1.id
    join (
        select
            max(g.ball),
            c.name
        from
            student_curs as s1
            join student as s2 on s1.student_id = s2.id
            join curs as c on curs_id = c.id
            join gred as g on g.student_curs_id = s1.id
        group by
            c
    ) as q on q.max = g.ball
    and q.name = c.name
group by
    c.name;

-- 3-savol
select
    c.name,
    round(avg(g.ball))
from
    student_curs as s1
    join student as s2 on s1.student_id = s2.id
    join curs as c on curs_id = c.id
    join gred as g on g.student_curs_id = s1.id
group by
    c.name;

-- 4-savols
select
    c.name,
    array_agg (s2.name) as name,
    array_agg (s2.age)
from
    student_curs as s1
    join student as s2 on s1.student_id = s2.id
    join curs as c on curs_id = c.id
    join gred as g on g.student_curs_id = s1.id
    join (
        select
            min(s2.age),
            c.name
        from
            student_curs as s1
            join student as s2 on s1.student_id = s2.id
            join curs as c on curs_id = c.id
            join gred as g on g.student_curs_id = s1.id
        group by
            c.name
    ) as q on q.min = s2.age
    and q.name = c.name
group by
    c.name;

-- 5-misol
select
    c.name,
    array_agg (g.ball)
from
    student_curs as s1
    join student as s2 on s1.student_id = s2.id
    join curs as c on curs_id = c.id
    join gred as g on g.student_curs_id = s1.id
where
    g.ball in (
        select
            g.ball
        from
            student_curs as s1
            join student as s2 on s1.student_id = s2.id
            join curs as c on curs_id = c.id
            join gred as g on g.student_curs_id = s1.id
        order by
            g.ball desc
        limit
            1
    )
group by
    c.name;