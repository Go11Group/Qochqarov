delete from 
    course 
        where id in 
            (select c.id from 
            course as c join 
            person as p 
            on c.person_id=p.id)
