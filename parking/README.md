Problem Statement: 

Design a parking lot that allows drivers to:
- Enter a lot and receive and assigned spot upon entry
- Exit a lot and have their spot be considered vacant upon exit
- See how many available spots are in the lot of varying sizes (big, medium, small)

/*
<>

INFORMATION: 

QUESTIONS & ASSUMPTIONS:
- 

CONSTRAINTS: 

ENTITIES:

- Vehicle
    - It has size and id

- Lot
    - It has size and id

- Parking
    - It has capacity of small lots
    - It has capacity of medium lots
    - It has capacity of big lots
    - It has map of cars on small lots
    - It has map of cars on medium lots
    - It has map of cars on big lots

    - It is able to locate a car in a lot when entering
    - It is able to remove car from a lot when exiting

COMPONENTS & INTERFACES: 
- We'll have a Vehicle component that mainly identifies a car size and id.
    - Vehicle Interface: GetID(), GetSize()
    - Constructor.

- We'll have a Lot component that mainly identifies a lot id and size of the lot and it's availability
    - Lot Interface: GetID(), GetSize(), SetAvailable(), SetUnavailable()
    - Constructor

- We'll have a Parking component that will contain an amount of lots of different sizes
    - Through this component we'll be able to register and remove cars from lots, given their size and lot availability.
    - Parking Interface: Park(Vehicle), Unpark(Vehicle)
    - Constructor(): Receives lots amount of each type and initializes the map.
        - The ids will be defined using offset: 
            - If small are 3, medium are 2 and big are 1
                - Small ids will be 1 to 3, medium ids will be from len(small) and big will be from len(small) + len(medium) 
    - Internal Logic: 
        - Given a car, check if we having available lots of that size
            - If we have, place it, update the availability of that lot and update the map of cars in lots
            - If we don't have, return err


<>
*/
