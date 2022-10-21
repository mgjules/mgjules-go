CREATE MIGRATION m1vfglkwpnn6xc7vfk2dtcswon2mrzqpptoqoyin7445mm25atpqeq
    ONTO m1mr3oizjj2nar7n33pllig34scuzus2fwo6wb3zi7563qfsunavca
{
  ALTER TYPE default::CVExperience {
      ALTER LINK technologies {
          ALTER PROPERTY sort {
              RESET default;
          };
      };
  };
};
