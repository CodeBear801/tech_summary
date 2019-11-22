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
import random
from operator import add

from pyspark.sql import SparkSession


if __name__ == "__main__":

    spark = SparkSession\
        .builder\
        .appName("hotspot")\
        .getOrCreate()

    sc = spark.sparkContext

    #list have hot spot, there are lots of 'a'
    skew_list=[('a',random.randint(1,1000)) for i in range(1000000)]
    skew_list.append(('b',10))
    skew_list.append(('c',8))
    #normal data list
    normal_list=[('a',9),('b',3),('c',8)]
    
    skew_rdd=sc.parallelize(skew_list)
    normal_rdd=sc.parallelize(normal_list)
    
    #1,sampling on skew_rdd，assume there is only one key has hot spot
    # the target is find that key
    skew_sample=skew_rdd.sample(False, 0.3, 9).groupByKey().mapValues(len)
    skew_sample.cache()
    #print (skew_sample.collect())
    #skew_sample.foreach(print)

    # [perry] sort is slow, we just need to find max element
    max_count_key, sample_max_count = sorted(skew_sample.collect(), key = lambda x:x[1], reverse = True)[0]
    print (max_count_key)
    #skew_sample_count_map=skew_sample.map(lambda (k,v):(len(v),k))
    #skew_sample_count_map.cache()
    #max_count=skew_sample_count_map.reduce(lambda x,y:max(x[0],y[0]))[0]
    #max_count_key=skew_sample_count_map.filter(lambda x:x[0]==max_count).collect()[0][1]
    
    #2，based on the max_count_key, divide skew_rdd to two parts, one contains max_count_key and the other don't
    # then apply join() operation with normal rdd, finally do union()
    max_key_rdd=skew_rdd.filter(lambda x:x[0]==max_count_key)
    other_key_rdd=skew_rdd.filter(lambda x:x[0]!=max_count_key)
    result1=max_key_rdd.join(normal_rdd) 
    result2=other_key_rdd.join(normal_rdd) 
    print (result1.union(result2).count())

    spark.stop()
