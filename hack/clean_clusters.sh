array=()
for i in "${array[@]}"
do
    ocm delete /api/clusters_mgmt/v1/clusters/$i
done
