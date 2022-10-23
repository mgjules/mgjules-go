CREATE MIGRATION m1mr3oizjj2nar7n33pllig34scuzus2fwo6wb3zi7563qfsunavca
    ONTO m1ec7wgoq6fnxmeo6qhqdrofekko2ttgh435kxvnn2f3jqcvh7bb7a
{
  ALTER TYPE default::CVExperience {
      ALTER LINK technologies {
          CREATE CONSTRAINT std::expression ON ((__subject__@sort >= 0));
      };
  };
};
