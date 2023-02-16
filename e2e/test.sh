echo "Iniciando Test E2E...."
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
echo "Migrando Database..."
./bin/mantenimiento migrate
echo ""
echo "Usuarios"
./bin/mantenimiento i -to usuarios -d '{"nombre": "luis", "email": "luis@me.com", "password": "Qwerty123"}'
./bin/mantenimiento ls -f usuarios
echo ""
echo "Empresas"
./bin/mantenimiento i -to empresas -d '{"nombre": "flander y asociados", "email": "luis@me.com", "direccion": "Ave. Siempre viva", "provincia Springfield", "owner": "luis@me.com"}'
./bin/mantenimiento ls -f empresas
echo ""
echo "Contactos"
./bin/mantenimiento i -to contactos -d '{"empresa_id": 1, "numero": "8299833824", "tipo": "CASA", "owner": "luis@me.com"}'
./bin/mantenimiento ls -f contactos
echo ""
echo "Geopocisiones"
./bin/mantenimiento i -to geopoints -d '{"empresa_id": 1, "lat": 87.0000, "lng": 78.0000, "owner": "luis@me.com"}'
./bin/mantenimiento ls -f geopoints
echo ""
echo "Areas"
./bin/mantenimiento i -to areas -d '{"empresa_id": 1, "descripcion": "Descripcion del area", "capacidad": 1000, "largo": 300, "ancho": 100, "owner": "luis@me.com"}'
./bin/mantenimiento ls -f areas
echo ""
echo "Tipos De Areas"
./bin/mantenimiento i -to areas_types -d '{"area_id": 1, "nombre": "nombre del area", "owner": "luis@me.com"}'
./bin/mantenimiento ls -f areas_types
echo ""
echo "Equipos"
./bin/mantenimiento i -to equipos -d '{"area_id": 1, "marca": "marca del equipo", "modelo": "modelo del equipo", "serie": "Serie del equipo", "descripcion": "descripcion del equipo", "owner": "luis@me.com"}'
./bin/mantenimiento ls -f equipos
echo ""
echo "Tipos De Equipos"
./bin/mantenimiento i -to equipos_types -d '{"equipo_id": 1, "nombre": "nombre del equipo", "owner": "luis@me.com"}'
./bin/mantenimiento ls -f equipos_types
echo ""
echo "Componentes"
./bin/mantenimiento i -to componentes -d '{"equipo_id": 1, "marca": "marca del componente", "modelo": "modelo del componente", "serie": "Serie del componente", "descripcion": "descripcion del componente", "vida_util": 1000, "owner": "luis@me.com"}'
./bin/mantenimiento ls -f componentes
echo ""
echo "Tipos De Componentes"
./bin/mantenimiento i -to componentes_types -d '{"componente_id": 1, "nombre": "nombre del componente", "owner": "luis@me.com"}'
./bin/mantenimiento ls -f componentes_types
echo ""
#echo "Iniciando Servidor en puerto 8080"
#./bin/mantenimiento s -p 8080
#echo ""
echo "Borrando datos..."
echo ""
./bin/mantenimiento clr -f componentes_types
echo ""
./bin/mantenimiento clr -f areas_types
echo ""
./bin/mantenimiento clr -f equipos_types
echo ""
./bin/mantenimiento clr -f contactos
echo ""
./bin/mantenimiento clr -f geopoints
echo ""
./bin/mantenimiento clr -f empresas
echo ""
./bin/mantenimiento clr -f areas
echo ""
./bin/mantenimiento clr -f equipos
echo ""
./bin/mantenimiento clr -f componentes
echo ""
./bin/mantenimiento clr -f usuarios
echo ""
echo "Apagandp Docker"
make down