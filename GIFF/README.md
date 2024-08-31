## Test

```
python test.py --data data/owndata.yaml --device 0 --weights best.pt
```

## Training

``` shell
python train.py --data data/owndata.yaml --cfg cfg/training/yolov7.yaml
```

## Inference

``` shell
python detect.py --weights best.pt --source DataSet\images\test
```



