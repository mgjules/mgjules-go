module default {
  scalar type Lang extending enum<En, Fr>;

  scalar type Gender extending enum<Male, Female, Robot>;

  scalar type Role extending enum<'Founder', 'Co-Founder', 'Speaker', 'Co-Speaker'>;

  abstract type Timestamps {
    required property created_at -> datetime {
      default := datetime_current();
      readonly := true;
    }
  }

  abstract type Avatar {
    required property avatar -> str{
      default := "https://mgjules.dev/img/avatar.webp";
    }
  }

  abstract type Sort {
    required property sort -> int32 {
      default := 0;
    }
  }

  abstract type Period {
    required property from -> cal::local_date;
    property to -> cal::local_date;
  }

  abstract link link_with_sort {
    property sort -> int32;
    constraint expression on (
      __subject__@sort >= 0
    );
  }

  type Meta extending Timestamps, Avatar {
    required property base_url -> str;
    required property lang -> Lang;
    required property first_name -> str;
    required property last_name -> str;
    required property gender -> Gender;
    required property description -> str;
    required property keywords -> array<str>;
    required property github -> str;
    required property username -> str;
  }

  type SiteLink extending Timestamps, Sort {
    required property name -> str;
    required property url -> str; 
    property alternate_url -> str; 
    required property new_window -> bool {
      default := false;
    };
    required property icon -> str;
  }

  type Introduction extending Timestamps, Avatar {
    required property introduction -> str;
  }

  type CVSection extending Timestamps, Sort {
    required property name -> str {
      constraint exclusive;
    };
    required property icon -> str;
  }

  type CVTechnology extending Timestamps {
    required property name -> str {
      constraint exclusive;
    };
  }

  type CVExperience extending Timestamps, Period {
    required property company -> str;
    required property position -> str {
      default := "Backend Engineer"
    };
    required property link -> str;
    multi link technologies extending link_with_sort -> CVTechnology;
    required property tasks -> array<str>;
  }

  type CVProject extending Timestamps, Sort {
    required property name -> str {
      constraint exclusive;
    };
    required property link -> str;
    required property description -> str;
    multi link technologies extending link_with_sort -> CVTechnology;
  }

  type CVContribution extending Timestamps, Period {
    required property event -> str;
    required property role -> Role;
    required property title -> str;
    required property link -> str;
  }

  type CVAward extending Timestamps {
    required property event -> str;
    required property description -> str;
    required property date -> cal::local_date;
    required property link -> str;
    required property icon -> str;
    required property result -> str;
  }

  type CVInterest extending Timestamps, Sort {
    required property name -> str;
    required property image -> str;
  }

  type CVLanguage extending Timestamps, Sort {
    required property name -> str;
    required property icon -> str;
    required property level -> str;
  }

  type User extending Timestamps, Avatar {
    required property first_name -> str;
    required property last_name -> str;
  }

  type BlogTag extending Timestamps {
    required property name -> str;
    required property slug := re_replace(r'[^a-z]+', '_', str_lower(.name), flags := 'g');
    index on (.slug);
  }

  type BlogPost extending Timestamps {
    required property title -> str;
    required property slug := re_replace(r'[^a-z]+', '_', str_lower(.title), flags := 'g');
    required property summary -> str;
    required property cover_image -> str {
      default := "https://mgjules.dev/img/blog/modern-code-screen.webp";
    };
    required property content -> str;
    required multi link tags -> BlogTag;
    required link created_by -> User {
      constraint exclusive;
    };
    index on (.slug);
  }
};
