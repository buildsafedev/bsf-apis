
{
	description = "";
	
	inputs = {
		 nixpkgs-6d974716629d3bf8a31b519952c6a5754a41bd38.url = "github:nixos/nixpkgs/6d974716629d3bf8a31b519952c6a5754a41bd38";
		 nixpkgs-ac5c1886fd9fe49748d7ab80accc4c847481df14.url = "github:nixos/nixpkgs/ac5c1886fd9fe49748d7ab80accc4c847481df14";
		 nixpkgs-a731d0cb71c58f56895f71a5b02eda2962a46746.url = "github:nixos/nixpkgs/a731d0cb71c58f56895f71a5b02eda2962a46746";
		 nixpkgs-1ebb7d7bba2953a4223956cfb5f068b0095f84a7.url = "github:nixos/nixpkgs/1ebb7d7bba2953a4223956cfb5f068b0095f84a7";
			
		nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
		 gomod2nix.url = "github:nix-community/gomod2nix";
		gomod2nix.inputs.nixpkgs.follows = "nixpkgs";
		
		
		
		
		
		

		 
	};
	
	outputs = inputs@{ self, nixpkgs, 
	 gomod2nix, 
	
	
	
	
	 nixpkgs-6d974716629d3bf8a31b519952c6a5754a41bd38, 
	 nixpkgs-ac5c1886fd9fe49748d7ab80accc4c847481df14, 
	 nixpkgs-a731d0cb71c58f56895f71a5b02eda2962a46746, 
	 nixpkgs-1ebb7d7bba2953a4223956cfb5f068b0095f84a7, 
	 }: let
	  supportedSystems = [ "x86_64-linux" "aarch64-darwin" "x86_64-darwin" "aarch64-linux" ];
	  
	  
	  forEachSupportedSystem = f: nixpkgs.lib.genAttrs supportedSystems (system: f {
		inherit system;
		
		 nixpkgs-6d974716629d3bf8a31b519952c6a5754a41bd38-pkgs = import nixpkgs-6d974716629d3bf8a31b519952c6a5754a41bd38 { inherit system; };
		 nixpkgs-ac5c1886fd9fe49748d7ab80accc4c847481df14-pkgs = import nixpkgs-ac5c1886fd9fe49748d7ab80accc4c847481df14 { inherit system; };
		 nixpkgs-a731d0cb71c58f56895f71a5b02eda2962a46746-pkgs = import nixpkgs-a731d0cb71c58f56895f71a5b02eda2962a46746 { inherit system; };
		 nixpkgs-1ebb7d7bba2953a4223956cfb5f068b0095f84a7-pkgs = import nixpkgs-1ebb7d7bba2953a4223956cfb5f068b0095f84a7 { inherit system; };
		
		 buildGoApplication = gomod2nix.legacyPackages.${system}.buildGoApplication;
		pkgs = import nixpkgs { inherit system;  };
		
		
	  });
	in {
	
	packages = forEachSupportedSystem ({ pkgs,
		 buildGoApplication, 
		
		
		 nixpkgs-6d974716629d3bf8a31b519952c6a5754a41bd38-pkgs, 
		 nixpkgs-ac5c1886fd9fe49748d7ab80accc4c847481df14-pkgs, 
		 nixpkgs-a731d0cb71c58f56895f71a5b02eda2962a46746-pkgs, 
		 nixpkgs-1ebb7d7bba2953a4223956cfb5f068b0095f84a7-pkgs, 
		 ... }: {
		default = pkgs.callPackage ./default.nix {
			 inherit buildGoApplication;
			go = pkgs.go_1_22; 
			
			
			
		};
	  });
	
	  devShells = forEachSupportedSystem ({ pkgs, 
		 buildGoApplication, 
		
		
		 nixpkgs-6d974716629d3bf8a31b519952c6a5754a41bd38-pkgs, 
		 nixpkgs-ac5c1886fd9fe49748d7ab80accc4c847481df14-pkgs, 
		 nixpkgs-a731d0cb71c58f56895f71a5b02eda2962a46746-pkgs, 
		 nixpkgs-1ebb7d7bba2953a4223956cfb5f068b0095f84a7-pkgs, 
		 ... }: {
		devShell = pkgs.mkShell {
		  # The Nix packages provided in the environment
		  packages =  [
			nixpkgs-6d974716629d3bf8a31b519952c6a5754a41bd38-pkgs.api-linter  
			nixpkgs-6d974716629d3bf8a31b519952c6a5754a41bd38-pkgs.buf  
			nixpkgs-ac5c1886fd9fe49748d7ab80accc4c847481df14-pkgs.delve  
			nixpkgs-a731d0cb71c58f56895f71a5b02eda2962a46746-pkgs.go  
			nixpkgs-1ebb7d7bba2953a4223956cfb5f068b0095f84a7-pkgs.gotools  
			
		  ];
		};
	  });
	
	  runtimeEnvs = forEachSupportedSystem ({ pkgs,
		 buildGoApplication, 
		
		
		 nixpkgs-6d974716629d3bf8a31b519952c6a5754a41bd38-pkgs,  nixpkgs-ac5c1886fd9fe49748d7ab80accc4c847481df14-pkgs,  nixpkgs-a731d0cb71c58f56895f71a5b02eda2962a46746-pkgs,  nixpkgs-1ebb7d7bba2953a4223956cfb5f068b0095f84a7-pkgs,  ... }: {
		runtime = pkgs.buildEnv {
		  name = "runtimeenv";
		  paths = [ 
			nixpkgs-ac5c1886fd9fe49748d7ab80accc4c847481df14-pkgs.cacert   
			
		   ];
		};
	   });

	   devEnvs = forEachSupportedSystem ({ pkgs,
		 buildGoApplication, 
		
		
	    nixpkgs-6d974716629d3bf8a31b519952c6a5754a41bd38-pkgs,  nixpkgs-ac5c1886fd9fe49748d7ab80accc4c847481df14-pkgs,  nixpkgs-a731d0cb71c58f56895f71a5b02eda2962a46746-pkgs,  nixpkgs-1ebb7d7bba2953a4223956cfb5f068b0095f84a7-pkgs,  ... }: {
		development = pkgs.buildEnv {
		  name = "devenv";
		  paths = [ 
			nixpkgs-6d974716629d3bf8a31b519952c6a5754a41bd38-pkgs.api-linter  
			nixpkgs-6d974716629d3bf8a31b519952c6a5754a41bd38-pkgs.buf  
			nixpkgs-ac5c1886fd9fe49748d7ab80accc4c847481df14-pkgs.delve  
			nixpkgs-a731d0cb71c58f56895f71a5b02eda2962a46746-pkgs.go  
			nixpkgs-1ebb7d7bba2953a4223956cfb5f068b0095f84a7-pkgs.gotools  
			
		   ];
		};
	   });
       
	   
	   
	};
}
