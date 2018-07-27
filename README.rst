Golang config loader
========================================================

Usage
====

.. code-block:: golang

    type Config struct {
        Service  Service `json:"Service"`
    }

    //File Config
	var configPath = flag.String("config", "./config.json", "path of config file")
	flag.Parse()

	// Load the configuration file
    var config Config
	err := gonf.Load(*configPath, &config)
	if err != nil {
		log.Fatal("Error loading configs: ", err)
	}

    log.Println("My configs: ", config)


Tests
====

.. code-block:: console
    
    make tests