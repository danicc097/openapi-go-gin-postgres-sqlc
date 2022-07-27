-- name: GetPetById :one
select
  pet_id,
  color,
  metadata
from
  pets
left join animals using (animal_id)
where pet_id = @pet_id
limit 1;


