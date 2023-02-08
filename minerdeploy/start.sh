#! /bin/bash
sh -c 'cd /app; ./geth --datadir dogcoinnetwork --networkid 77612 --bootnodes "enr:-KO4QGBjjsdxCdrdkbgnGCcPMAJDGrHIhRdJgKLn_APqFa7VVAlS-Q55FSYvVaAGnwAboohy36O1QOlYdowjpTZUeS2GAYXEIc11g2V0aMfGhA_wvjOAgmlkgnY0gmlwhM-Ua4aJc2VjcDI1NmsxoQOF_sXeyV-fgeOk2LQTqbsUfv4WhktMKuE3bZ2APJETyYRzbmFwwIN0Y3CCf5yDdWRwgn-c" --mine --syncmode full --unlock 0 --password /app/.accountpassword'
