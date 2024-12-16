CREATE MIGRATION m1jm33r7e7rdbuog2nlkecqcspu44ew4lo4uaye7jpsov4ntchlr5q
    ONTO m1ipbocrwwz6trrjtxfoovgywto262h5d2mm3ejpzekevrgvkfjduq
{
  ALTER TYPE default::CVProject EXTENDING default::Sort LAST;
};
