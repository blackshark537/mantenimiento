echo "Iniciando...."
echo ""
echo "Ejecutando Docker Compose"
echo ""
make up
echo ""
echo "Compilando..."
make build
echo ""
echo "Iniciando CLI Test Commands..."
echo ""
echo "Usuarios"
./bin/mantenimiento ls -f users
echo ""
echo "Empresas"
./bin/mantenimiento ls -f empresas
echo ""
echo "Contactos"
./bin/mantenimiento ls -f contactos
echo ""
echo "Geopocisiones"
./bin/mantenimiento ls -f geopoints
echo ""
echo "Areas"
./bin/mantenimiento ls -f areas
echo ""
echo "Tipos De Areas"
./bin/mantenimiento ls -f areas_types
echo ""
echo "Equipos"
./bin/mantenimiento ls -f equipos
echo ""
echo "Tipos De Equipos"
./bin/mantenimiento ls -f equipos_types
echo ""
echo "Componentes"
./bin/mantenimiento ls -f componentes
echo ""
echo "Tipos De Componentes"
./bin/mantenimiento ls -f componentes_types
echo ""
echo "Iniciando Servidor en puerto 8080"
./bin/mantenimiento s -p 8080
echo ""
echo "Apagandp Docker"
make down