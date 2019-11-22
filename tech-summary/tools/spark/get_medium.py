#
# Licensed to the Apache Software Foundation (ASF) under one or more
# contributor license agreements.  See the NOTICE file distributed with
# this work for additional information regarding copyright ownership.
# The ASF licenses this file to You under the Apache License, Version 2.0
# (the "License"); you may not use this file except in compliance with
# the License.  You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

from __future__ import print_function

import sys
import array

from pyspark.sql import SparkSession


if __name__ == "__main__":

    spark = SparkSession\
        .builder\
        .appName("Medium")\
        .getOrCreate()
    sc = spark.sparkContext

    raw_arr = array.array('i',(i for i in range(0,1234)))
    #print (arr)
    sorted_array=sc.parallelize(raw_arr)
    sorted_array.cache()

    group_element_count=sorted_array.map(lambda e:(e//10,1)) \
                                    .reduceByKey(lambda x,y : x+y) \
                                    .sortByKey()
    #print(group_element_count.collect())
    group_element_count_map=group_element_count.collectAsMap()
    #print(group_element_count_map)
    element_count = 0
    for value in group_element_count_map.values():
        element_count += value

    temp = 0
    index = 0
    mid = 0
    if element_count%2!=0:
        mid=element_count//2+1
    else:
        mid=element_count//2
    for i in range(group_element_count.count()):
        temp += group_element_count_map[i]
        if temp>=mid:
            index=i
            break
    offset=mid - (temp - group_element_count_map[index])
    #print (index, offset, type(offset))
    
    result = sorted_array.map(lambda e : (e//10,e)) \
                          .filter(lambda k : k[0] == index) \
                          .takeOrdered(offset, key=lambda k:k[1])[-1][1]
                          
    print (result)

    spark.stop()
