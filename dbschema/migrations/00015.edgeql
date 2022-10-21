CREATE MIGRATION m1njgyyvp2gx5mqjrkxpcbbq6xpymu7xqliy2vel6acokyg4qovhea
    ONTO m1vfglkwpnn6xc7vfk2dtcswon2mrzqpptoqoyin7445mm25atpqeq
{
  ALTER TYPE default::CVExperience {
      CREATE REQUIRED PROPERTY position -> std::str {
          SET default := 'Backend Engineer';
      };
  };
};
